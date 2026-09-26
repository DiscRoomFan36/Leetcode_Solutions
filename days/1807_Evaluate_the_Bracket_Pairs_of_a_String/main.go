package main

import "strings"

func evaluate(s string, knowledge [][]string) string {
    knowledge_map := make(map[string]string);
    for _, know := range knowledge {
        key, value := know[0], know[1];
        knowledge_map[key] = value;
    }

    result := strings.Builder{};
    for {
        left, right, start_found := strings.Cut(s, "(");
        result.WriteString(left);

        if !start_found { break; }

        key, rest, end_found := strings.Cut(right, ")");
        if !end_found { panic("this input is not valid"); }

        value, in_map := knowledge_map[key];
        if in_map {
            result.WriteString(value);
        } else {
            result.WriteString("?");
        }

        s = rest;
    }

    return result.String();
}
