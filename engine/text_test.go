package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	assert := assert.New(t)

	t.Run("CountWord", func(t *testing.T) {
		text := "Contrairement à une opinion répandue, le Lorem Ipsum n'est pas simplement du texte aléatoire.\nIl trouve ses racines dans une oeuvre de la littérature latine classique datant de 45 av. J.-C., le rendant vieux de 2000 ans.\n\nUn professeur du Hampden-Sydney College, en Virginie, s'est intéressé à un des mots latins les plus obscurs, consectetur, extrait d'un passage du Lorem Ipsum, et en étudiant tous les usages de ce mot dans la littérature classique, découvrit la source incontestable du Lorem Ipsum."

		assert.Equal(3, CountWordInText(text, "la"), "should be equal")
	})

	t.Run("GetWordsFromText", func(t *testing.T) {
		text := "Un professeur du Hampden-Sydney College, en Virginie, s'est intéressé à un des mots latins les plus obscurs, consectetur, extrait d'un passage du Lorem Ipsum."

		assert.Equal([]string{"un", "professeur", "du", "hampden-sydney", "college", "en", "virginie", "s'est", "intéressé", "à", "un", "des", "mots", "latins", "les", "plus", "obscurs", "consectetur", "extrait", "d'un", "passage", "du", "lorem", "ipsum"}, GetWordsFromText(text), "should be equal")
	})

	t.Run("IsTextFile", func(t *testing.T) {
		assert.Equal(true, IsTextFile([]byte("Here is a string....")), "should be equal")

		assert.Equal(false, IsTextFile([]byte("<html></html>")), "should be equal")
	})

	t.Run("GetFirstMatchingLine", func(t *testing.T) {
		text := "1 no\n2no\n3 yes"

		assert.Equal("3 yes", GetFirstMatchingLine(text, "yes"), "should be equal")
	})
}
