package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fichier, _ := ioutil.ReadFile("capsule_file/Hangman.txt")                                 // peut varier selon le chemin du fichier. Récupère le fichier Hangman.txt à l'aide de ioutil et l'assigne à la variable fichier, "_" permet de ne pas récupérer l'erreur
	str := string(fichier)                                                                    // transforme la variable fichier de type []byte en chaine de caractère et l'assigne à str
	fmt.Println(premier_mot(str))                                                             // affiche le premier fragment
	fmt.Println("\n", dernier_mot(str))                                                       // affiche le deuxième fragment
	fmt.Println(element_index_str(element_apres(chaque_mot(str), "before"), chaque_mot(str))) // affiche le troisième fragment
	fmt.Println(dernier_fragment(chaque_mot(str), "now"))                                     // affiche le quatrième fragment
	fmt.Print("\n\n")                                                                         //
	consigne()                                                                                // affiche ce qui est demandé dans les 4 fragments
}

//____________________________________________	PREMIER FRAGMENT	______________________________________________________________

func premier_mot(text string) string {
	// revoie le premier mot du fichier passé en string
	mot := ""
	for i := 0; text[i] != '\n'; i++ {
		mot += string(text[i])
	}
	return mot
}

//____________________________________________	SECOND FRAGMENT___________________________________________________________________

func dernier_mot(text string) string {
	// revoie le dernier mot du fichier passé en string
	mot := ""
	for i := len(text) - 1; text[i] != '\n'; i-- {
		mot += string(text[i])
	}
	return mot
}

//____________________________________________	TROISIEME FRAGMENT_________________________________________________________________

func chaque_mot(text string) []string {
	// transforme la chaine de caractère obtenue en lisant le fichier en tableau de string
	liste := []string{}
	mot := ""
	for _, element := range text { // parcours de la string text
		if element == 13 { // vérifie si l'élément est un retour à la ligne
			liste = append(liste, mot) // si c'est le cas ajoute le mot au tableau
			mot = ""                   // et réinitialise la variable mot
		} else { // si l'élément n'est pas un retour à la ligne
			mot += string(element) // l'élément est ajouté au mot en attendant de rencontrer un retour à la ligne
		}
	}
	if mot != "" { // vérifie si il reste quelque chose dans la variable mot
		liste = append(liste, mot) // si mot n'est pas vide, ajout de ce qui reste dans mot au tableau
	}
	return liste
}

func element_avant(liste []string, mot string) string {
	var resultat string
	for i := 0; i < len(liste); i++ { // initialise un index i qui va rester strictement inférieur à len(liste) afin de parcourir chaque élément de liste
		if strings.Contains(liste[i], mot) { //vérifie si le mot spécifié est présent dans l'élément de liste
			resultat = string(liste[i-1]) // assigne l'élément i-1 de liste à resultat, cela récupère la string avant la première itération du mot dans liste
		}
	}
	return resultat
}

func element_apres(liste []string, mot string) string {
	// fait de même que la fonction précédente mais récupère la string qui suit le mot spécifié
	var resultat string
	for i := 0; i < len(liste); i++ {
		if strings.Contains(liste[i], mot) {
			resultat = string(liste[i+1])
		}
	}
	return resultat
}

func element_index_str(index string, liste []string) string {
	var resultat string
	var index_modif int
	copie := ""
	for _, element := range index { // parcours les éléments de la string index
		if element >= 48 && element <= 57 { // et ajoute l'élément uniquement si c'est un chiffre
			copie += string(element)
		}
	}

	index_modif, _ = strconv.Atoi(copie) // converti la chaine de caractère en entier
	index_modif -= 1                     // converti l'index de manière à  ce qu'il corresponde à celui d'un tableau (qui commence à l'index 0)
	resultat = liste[index_modif]        // assigne la valeur de l'élément d'indice "index_modif" de liste à la variable résultat
	return resultat
}

// ______________________________________________	DERNIER FRAGMENT	_____________________________________________________________

func dernier_fragment(liste []string, mot string) string {
	// réalise ce qui est demandé
	mot_avant_now := element_avant(liste, mot)           // récupère le mot présent avant le mot "now" (ou tout autre mot pris en paramètre)
	index_resultat := int(mot_avant_now[1]) / len(liste) //assigne le calcul suivant à la variable "index_resultat" : second byte du mot précédant now / nombre de mots ou nombres dans le fichier
	resultat := liste[index_resultat]                    // assigne l'élément de liste d'indice "index_resultat" à la variable resultat
	return resultat
}

//______________________________________	CONSIGNE DONNEE PAR LES 4 FRAGMENT	______________________________________________________

func consigne() {
	// affiche un entier aléatoire en utilisant time pour mettre à jour la seed de rand à chaque appel
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100) // génère un entier pseudo-aléatoire entre 0 et 100
	fmt.Println("Voici un entier aléatoire compris entre 0 et 100 généré avec math/rand et time :", random)
}
