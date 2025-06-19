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

	// Check that zones field exists and has the expected structure
	zones, ok := result["zones"]
	if !ok {
		t.Fatal("JSON output missing 'zones' field")
	}

	// For --tags brand, expect: {"zones": {"tags": {"brand": [...]}}}
	zonesObj, ok := zones.(map[string]interface{})
	if !ok {
		t.Fatal("'zones' field is not an object")
	}

	tags, ok := zonesObj["tags"]
	if !ok {
		t.Fatal("'zones.tags' field missing")
	}

	tagsObj, ok := tags.(map[string]interface{})
	if !ok {
		t.Fatal("'zones.tags' field is not an object")
	}

	brandZones, ok := tagsObj["brand"]
	if !ok {
		t.Fatal("'zones.tags.brand' field missing")
	}

	zonesArray, ok := brandZones.([]interface{})
	if !ok {
		t.Fatal("'zones.tags.brand' field is not an array")
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

func TestJSONOutputMultipleTags(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "zonedb_test", "./cmd/zonedb")
	cmd.Dir = "../../"
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build zonedb: %v", err)
	}
	defer os.Remove("../../zonedb_test")

	// Test JSON output for multiple tags
	cmd = exec.Command("./zonedb_test", "--tags", "brand,generic", "--json")
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

	// For multiple tags, expect: {"zones": {"tags": {"all_of": [...], "domains": [...]}}}
	zones, ok := result["zones"]
	if !ok {
		t.Fatal("JSON output missing 'zones' field")
	}

	zonesObj, ok := zones.(map[string]interface{})
	if !ok {
		t.Fatal("'zones' field is not an object")
	}

	tags, ok := zonesObj["tags"]
	if !ok {
		t.Fatal("'zones.tags' field missing")
	}

	tagsObj, ok := tags.(map[string]interface{})
	if !ok {
		t.Fatal("'zones.tags' field is not an object")
	}

	allOf, ok := tagsObj["all_of"]
	if !ok {
		t.Fatal("'zones.tags.all_of' field missing")
	}

	allOfArray, ok := allOf.([]interface{})
	if !ok {
		t.Fatal("'zones.tags.all_of' field is not an array")
	}

	// Check that we have the expected tags
	if len(allOfArray) != 2 {
		t.Fatalf("Expected 2 tags, got %d", len(allOfArray))
	}

	// Check for brand and generic tags
	foundBrand, foundGeneric := false, false
	for _, tag := range allOfArray {
		tagStr, ok := tag.(string)
		if !ok {
			t.Fatal("Tag is not a string")
		}
		if tagStr == "brand" {
			foundBrand = true
		} else if tagStr == "generic" {
			foundGeneric = true
		}
	}

	if !foundBrand {
		t.Fatal("Expected 'brand' tag not found")
	}
	if !foundGeneric {
		t.Fatal("Expected 'generic' tag not found")
	}

	// Check filtered array exists
	filtered, ok := tagsObj["filtered"]
	if !ok {
		t.Fatal("'zones.tags.filtered' field missing")
	}

	filteredArray, ok := filtered.([]interface{})
	if !ok {
		t.Fatal("'zones.tags.filtered' field is not an array")
	}

	if len(filteredArray) == 0 {
		t.Fatal("No filtered zones found in JSON output")
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
