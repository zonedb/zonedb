package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestJSONOutput(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "zonedb_test", "./cmd/zonedb")
	cmd.Dir = "../../"
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build zonedb: %v", err)
	}
	defer os.Remove("../../zonedb_test")

	// Test JSON output for zones
	cmd = exec.Command("./zonedb_test", "--tags", "brand", "--json")
	cmd.Dir = "../../"
	cmd.Env = append(os.Environ(), "FORCE_COLOR=0") // Disable colors for testing
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run zonedb with JSON flag: %v", err)
	}

	// Parse JSON output
	var result map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Check that zones field exists and is an array
	zones, ok := result["zones"]
	if !ok {
		t.Fatal("JSON output missing 'zones' field")
	}

	zonesArray, ok := zones.([]interface{})
	if !ok {
		t.Fatal("'zones' field is not an array")
	}

	// Check that we have some zones
	if len(zonesArray) == 0 {
		t.Fatal("No zones found in JSON output")
	}

	// Check that the zones are strings
	for i, zone := range zonesArray {
		if _, ok := zone.(string); !ok {
			t.Fatalf("Zone at index %d is not a string: %v", i, zone)
		}
	}

	// Test that "aaa" is in the list (should be present for brand tag)
	found := false
	for _, zone := range zonesArray {
		if zone.(string) == "aaa" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Expected zone 'aaa' not found in brand zones")
	}
}

func TestJSONOutputTags(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "zonedb_test", "./cmd/zonedb")
	cmd.Dir = "../../"
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build zonedb: %v", err)
	}
	defer os.Remove("../../zonedb_test")

	// Test JSON output for tags
	cmd = exec.Command("./zonedb_test", "--list-tags", "--json")
	cmd.Dir = "../../"
	cmd.Env = append(os.Environ(), "FORCE_COLOR=0") // Disable colors for testing
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run zonedb with JSON flag: %v", err)
	}

	// Parse JSON output
	var result map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Check that tags field exists and is an array
	tags, ok := result["tags"]
	if !ok {
		t.Fatal("JSON output missing 'tags' field")
	}

	tagsArray, ok := tags.([]interface{})
	if !ok {
		t.Fatal("'tags' field is not an array")
	}

	// Check that we have some tags
	if len(tagsArray) == 0 {
		t.Fatal("No tags found in JSON output")
	}

	// Check that the tags are strings
	for i, tag := range tagsArray {
		if _, ok := tag.(string); !ok {
			t.Fatalf("Tag at index %d is not a string: %v", i, tag)
		}
	}

	// Test that "brand" is in the list
	found := false
	for _, tag := range tagsArray {
		if tag.(string) == "brand" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Expected tag 'brand' not found in tags list")
	}
}

func TestNonJSONOutputStillWorks(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "zonedb_test", "./cmd/zonedb")
	cmd.Dir = "../../"
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build zonedb: %v", err)
	}
	defer os.Remove("../../zonedb_test")

	// Test non-JSON output still works
	cmd = exec.Command("./zonedb_test", "--tags", "brand")
	cmd.Dir = "../../"
	cmd.Env = append(os.Environ(), "FORCE_COLOR=0") // Disable colors for testing
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run zonedb without JSON flag: %v", err)
	}

	// Check that output contains "Zones:" (non-JSON format)
	outputStr := string(output)
	if !strings.Contains(outputStr, "Zones:") {
		t.Fatal("Non-JSON output missing 'Zones:' prefix")
	}

	// Check that output contains some expected zones
	if !strings.Contains(outputStr, "aaa") {
		t.Fatal("Non-JSON output missing 'aaa' zone")
	}
}
