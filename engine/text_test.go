package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	assert := assert.New(t)

	t.Run("CountWord", func(t *testing.T) {
		t.Run("should count words in text", func(t *testing.T) {
			text := "Contrairement à une opinion répandue, le Lorem Ipsum n'est pas simplement du texte aléatoire. Il trouve ses racines dans une oeuvre de la littérature latine classique datant de 45 av. J.-C., le rendant vieux de 2000 ans. Un professeur du Hampden-Sydney College, en Virginie, s'est intéressé à un des mots latins les plus obscurs, consectetur, extrait d'un passage du Lorem Ipsum, et en étudiant tous les usages de ce mot dans la littérature classique, découvrit la source incontestable du Lorem Ipsum."

			assert.Equal(3, CountWord(text, "la"), "should be equal")
		})
	})
}
