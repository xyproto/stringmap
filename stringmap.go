package stringmap

import (
	"strings"
	"bytes"
	"sort"
)

type StringMap map[string]string

// Takes a list of strings in this format: "a:b", "b:c", "d:e" or
// a string in this format: "a:b, b:c, d:e" and creates a StringMap
func New(keyValueStrings... string) *StringMap {
	sm := make(StringMap)
	if (len(keyValueStrings) == 1) && strings.Contains(keyValueStrings[0], ",") {
		// Assume the first string is a comma-separated list
		keyValueStrings = strings.Split(keyValueStrings[0], ",")
	}
	var fields []string
	var key, value string
	for _, value = range keyValueStrings {
		fields = strings.SplitN(value, ":", 2)
		if len(fields) == 2 {
			key = strings.TrimSpace(fields[0])
			value = strings.TrimSpace(fields[1])
			if (key != "") && (value != "") {
				sm[key] = value
			}
		}
	}
	return &sm
}

// Convert a StringMap to a slice of strings where the keys and values are colon-separated
func (sm *StringMap) Slice() []string {
	s := make([]string, 0, len(*sm))
	for key, value := range *sm {
		s = append(s, key + ":" + value)
	}
	return s
}

// Return a comma-separated string that represents the StringMap
func (sm *StringMap) String() string {
	var buffer bytes.Buffer
	for _, value := range (*sm).Slice() {
		buffer.WriteString(value + ", ")
	}
	s := buffer.String()
	if s != "" {
		// Remove the last ", " at the end
		return s[:len(s)-2]
	}
	return s
}

// Map a string function over the values of the StringMap
func (sm *StringMap) Map(fn func (string) string) {
	for key, value := range *sm {
		// Assign a new value
		(*sm)[key] = fn(value)
	}
}

// Map a string function over the keys of the StringMap
func (sm *StringMap) MapKeys(fn func (string) string) {
	for key, value := range *sm {
		// Remove the old key
		delete(*sm, key)
		// Assign with the new key
		(*sm)[fn(key)] = value
	}
}

// Return the keys of the StringMap
func (sm *StringMap) Keys() []string {
	keys := make([]string, 0, len(*sm))
	for key, _ := range *sm {
		keys = append(keys, key)
	}
	return keys
}

// Return the values of the StringMap
func (sm *StringMap) Values() []string {
	values := make([]string, 0, len(*sm))
	for _, value := range *sm {
		values = append(values, value)
	}
	return values
}

// Return the sorted keys
func (sm *StringMap) SortedKeys() []string {
	keys := sm.Keys()
	sort.Strings(keys)
	return keys
}

// Check if a slice of strings has a given string
func has(sl []string, s string) bool {
	for _, element := range sl {
		if s == element {
			return true
		}
	}
	return false
}

// Given a slice of strings, find a string, return the position or -1
func find(sl []string, s string) int {
	for pos, element := range sl {
		if s == element {
			return pos
		}
	}
	return -1
}

// Return the keys, sorted by the values
func (sm *StringMap) KeysSortedByValues() []string {
	// 1. Get the values in sorted order
	sortedValues := sm.SortedValues()

	retkeys := make([]string, len(*sm))
	index := 0
	// 2. Go through all the keys and values
	for key, value := range *sm {
		// Find the index in sortedValues
		index = find(sortedValues, value)
		// Put the key in the right place in the keys to return
		if index != -1 {
			retkeys[index] = key
		}
	}

	// Return the sorted keys
	return retkeys
}

// Return the values in sorted order
func (sm *StringMap) SortedValues() []string {
	values := sm.Values()
	sort.Strings(values)
	return values
}

// Return as a map[string]string
func (sm *StringMap) GetMap() map[string]string {
	return map[string]string(*sm)
}

