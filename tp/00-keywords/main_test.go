package main

import (
	"go/scanner"
	"go/token"
	"os"
	"reflect"
	"testing"
)

// Liste des mots-clés Go à couvrir
var goKeywords = map[token.Token]bool{
	token.BREAK:       true,
	token.DEFAULT:     true,
	token.FUNC:        true,
	token.INTERFACE:   true,
	token.SELECT:      true,
	token.CASE:        true,
	token.DEFER:       true,
	token.GO:          true,
	token.MAP:         true,
	token.STRUCT:      true,
	token.CHAN:        true,
	token.ELSE:        true,
	token.GOTO:        true,
	token.PACKAGE:     true,
	token.SWITCH:      true,
	token.CONST:       true,
	token.FALLTHROUGH: true,
	token.IF:          true,
	token.RANGE:       true,
	token.TYPE:        true,
	token.CONTINUE:    true,
	token.FOR:         true,
	token.IMPORT:      true,
	token.RETURN:      true,
	token.VAR:         true,
}

// CountKeywordsUsingLexer analyse le code source et compte les occurrences des mots-clés Go
func CountKeywordsUsingLexer(code string) (map[string]int, map[string]bool) {
	// Initialisation de la table de comptage des mots-clés
	keywordCounts := make(map[string]int)
	keywordFound := make(map[string]bool)

	// Initialisation des positions et du scanner
	src := []byte(code)
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))

	// Scanner pour parcourir chaque token du code
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	// Boucle pour parcourir chaque token du code
	for {
		_, tok, _ := s.Scan()
		if tok == token.EOF {
			break // Arrêt à la fin du fichier
		}

		// Vérifie si le token est un mot-clé Go et l'incrémente dans les compteurs
		if goKeywords[tok] {
			keywordCounts[tok.String()]++
			keywordFound[tok.String()] = true
		}
	}

	return keywordCounts, keywordFound
}

func TestCountKeywordsUsingLexer(t *testing.T) {
	code, err := os.ReadFile("main.go")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	expected := map[string]bool{
		"package":     true,
		"import":      true,
		"func":        true,
		"var":         true,
		"for":         true,
		"if":          true,
		"continue":    true,
		"else":        true,
		"return":      true,
		"switch":      true,
		"case":        true,
		"fallthrough": true,
		"default":     true,
		"defer":       true,
		"map":         true,
		"struct":      true,
		"chan":        true,
		"go":          true,
		"select":      true,
		"goto":        true,
		"range":       true,
		"type":        true,
		"const":       true,
		"interface":   true,
	}

	resultNb, resultFound := CountKeywordsUsingLexer(string(code))

	if !reflect.DeepEqual(resultFound, expected) {
		t.Errorf("CountKeywordsUsingLexer() = %v; want %v", resultFound, expected)
	}
	scrore := 0
	for k, v := range resultNb {
		if expected[k] {
			scrore += v
		}
	}
	t.Log("Score:", scrore)
}
