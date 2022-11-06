package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

func exitOnError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "\n", err)
		os.Exit(1)
	}
}

func stateDir() string {
	workingDir, err := os.Getwd()
	exitOnError("Could not get working directory", err)
	return fmt.Sprintf("%s/state", workingDir)
}

func prepare() {
	// local backend requires a passphrase
	err := os.Setenv("PULUMI_CONFIG_PASSPHRASE", "test")
	exitOnError("Could not set PULUMI_CONFIG_PASSPHRASE", err)
	stateDir := stateDir()
	// make state dir
	err = os.Mkdir(stateDir, 0755)
	exitOnError("Could not create state directory", err)
	_, err = exec.Command("pulumi", "login", "--cloud-url", fmt.Sprintf("file://%s", stateDir)).Output()
	_, err = exec.Command("pulumi", "stack", "init", "dev").Output()
	exitOnError("Could not init pulumi stack", err)
}

func cleanup() {
	_, err := exec.Command("pulumi", "logout").Output()
	exitOnError("Could not logout from pulumi", err)
	stateDir := stateDir()
	err = os.RemoveAll(stateDir)
	exitOnError("Could not remove state directory", err)
	err = os.RemoveAll("Pulumi.dev.yaml")
	exitOnError("Could not remove Pulumi.dev.yaml", err)
}

func outputs() map[string]interface{} {
	_, err := exec.Command("pulumi", "up", "--stack", "dev", "--yes").Output()
	exitOnError("Could not run pulumi up", err)
	stdout, err := exec.Command("pulumi", "stack", "output", "--json").Output()
	exitOnError("Could not get pulumi stack outputs", err)
	var outputs map[string]interface{}
	err = json.Unmarshal(stdout, &outputs)
	exitOnError("Could not unmarshal pulumi stack outputs", err)
	return outputs
}

func jsonDeepEquals(a, b interface{}) bool {
	aBytes, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bBytes, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return string(aBytes) == string(bBytes)
}

func expectedOutputs() map[string]interface{} {
	return map[string]interface{}{
		"abs-positive":         42.24,
		"abs-negative":         42.24,
		"base64encode":         "QUJDREU=",
		"base64decode":         "ABCDE",
		"basename":             "pulumi",
		"base64gzip":           "H4sIAAAAAAAA/0pMSgYAAAD//wEAAP//wkEkNQMAAAA=",
		"base64sha256":         "ungWv48Bz+pBQUDeXa4iI7ADYaOWF3qctBD/YfIAFa0=",
		"base64sha512":         "3a81oZNherrMQXNJriBBMRLm+k6JqX6iCp7u5ktV05ohkpkqJ0/BqDa6PCOj/uu9RU1EI2Q86A4qmslPpUyknw==",
		"chunklist":            []interface{}{[]int{1, 2}, []int{3, 4}, []int{5}},
		"chunklistByOne":       []interface{}{[]int{1}, []int{2}, []int{3}, []int{4}, []int{5}},
		"cidrhost":             "10.0.0.2",
		"cidrhostV2":           "10.255.255.254",
		"cidrnetmask":          "255.0.0.0",
		"cidrsubnet":           "10.2.0.0/16",
		"coalesce":             "hello",
		"coalescelist":         []interface{}{1},
		"compact":              []string{"one"},
		"concat":               []int{1, 2, 3, 4, 5},
		"containsWithIntegers": true,
		"containsWithStrings":  true,
		"containsWithLists":    true,
		"shouldNotContain":     false,
		"dirname":              "/path/to/directory",
		"distinctWithIntegers": []int{1, 2, 3, 4, 5},
		"distinctWithStrings":  []string{"one", "two", "three"},
		"element":              20,
		"elementOverflow":      10,
		"elementNegativeIndex": 30,
	}
}

func main() {
	cleanup()
	prepare()
	defer cleanup()

	expected := expectedOutputs()
	outputs := outputs()
	for key, value := range outputs {
		foundValue, ok := expected[key]
		if !ok {
			fmt.Printf("💡 Unexpected output with key '%s'\n", key)
			continue
		}

		if reflect.DeepEqual(foundValue, value) || jsonDeepEquals(foundValue, value) {
			fmt.Printf("✅ Output '%s' has value '%v' as expected\n", key, value)
		} else {
			fmt.Printf("❌ Output '%s' was '%v' but should be '%v'\n", key, value, foundValue)
			return
		}
	}
}
