
package services

import (
	"math/rand"
	"context"
)

type Message struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

var messages = []string{
	"Bonjour !",
	"Salut, comment ça va ?",
	"Bienvenue sur notre API.",
	"Passez une excellente journée !",
	"Voici un message aléatoire.",
	"Le Go, c'est génial !",
	"Profitez de votre café.",
	"Restez motivé !",
	"Apprenez quelque chose de nouveau aujourd'hui.",
	"Prenez une pause.",
	"Vous êtes incroyable !",
	"Continuez comme ça.",
	"Le succès arrive à ceux qui persévèrent.",
	"Un pas à la fois.",
	"Gardez le sourire.",
	"Chaque jour compte.",
	"Soyez curieux.",
	"Partagez la bonne humeur.",
	"Merci de votre visite.",
	"À bientôt sur notre API !",
	"Rêvez grand.",
	"Prenez soin de vous.",
	"Chaque effort compte.",
	"Restez positif.",
	"Faites de votre mieux.",
	"Le savoir est une force.",
	"Soyez vous-même.",
	"Croyez en vos rêves.",
	"Le changement commence par vous.",
	"Ne baissez jamais les bras.",
	"Faites une pause et respirez.",
	"Le bonheur est contagieux.",
	"Soyez reconnaissant.",
	"Apprenez de vos erreurs.",
	"Partagez vos connaissances.",
	"Restez humble.",
	"Faites confiance au processus.",
	"Le temps est précieux.",
	"Soyez attentif aux détails.",
	"Exprimez votre créativité.",
	"Prenez des initiatives.",
	"Restez concentré.",
	"Faites preuve de gentillesse.",
	"Écoutez avant de parler.",
	"Travaillez en équipe.",
	"Valorisez la diversité.",
	"Acceptez les défis.",
	"Faites preuve de résilience.",
	"Restez curieux.",
	"Explorez de nouvelles idées.",
	"Soyez ouvert d'esprit.",
	"Prenez des risques calculés.",
	"Faites preuve de patience.",
	"Restez organisé.",
	"Planifiez votre journée.",
	"Fixez-vous des objectifs.",
	"Faites des pauses régulières.",
	"Riez souvent.",
	"Soyez flexible.",
	"Appréciez les petites choses.",
	"Faites preuve d'empathie.",
	"Soyez un bon leader.",
	"Encouragez les autres.",
	"Partagez vos réussites.",
	"Apprenez chaque jour.",
	"Restez motivé malgré les obstacles.",
	"Faites preuve de gratitude.",
	"Prenez le temps de réfléchir.",
	"Soyez proactif.",
	"Restez authentique.",
	"Faites confiance à votre intuition.",
	"Acceptez les critiques constructives.",
	"Restez à l'écoute.",
	"Faites preuve de persévérance.",
	"Soyez ponctuel.",
	"Respectez les autres.",
	"Faites preuve de générosité.",
	"Restez humble dans la réussite.",
	"Apprenez à dire non.",
	"Prenez soin de votre santé.",
	"Faites du sport régulièrement.",
	"Équilibrez travail et vie personnelle.",
	"Soyez fier de vos progrès.",
	"Faites preuve de curiosité.",
	"Restez à jour dans vos connaissances.",
	"Partagez vos idées.",
	"Faites preuve de rigueur.",
	"Soyez un modèle pour les autres.",
	"Restez fidèle à vos valeurs.",
	"Faites preuve de tolérance.",
	"Acceptez l'échec comme une leçon.",
	"Restez calme sous pression.",
	"Faites preuve de détermination.",
	"Soyez persévérant.",
	"Faites preuve de créativité.",
	"Restez humble face au succès.",
	"Soyez attentif à votre environnement.",
	"Faites preuve de bienveillance.",
	"Restez optimiste.",
	"Faites preuve de discipline.",
	"Soyez un apprenant à vie.",
	"Faites preuve de respect.",
	"Restez passionné.",
	"Faites preuve d'intégrité.",
	"Soyez inspirant.",
	"Restez engagé.",
	"Faites preuve de solidarité.",
	"Soyez innovant.",
}

func GetRandomMessage(ctx context.Context) Message {

	// Exemple d'utilisation du contexte (ici, non utilisé, mais prêt pour une base de données ou une annulation)
	id := rand.Intn(len(messages))
	return Message{
		Id:      id,
		Message: messages[id],
	}

}

func GetAllMessages() []Message {
	var ms []Message
	for i, message := range messages {
		ms = append(ms, Message{
			Id:      i,
			Message: message,
		})
	}
	return ms
}


