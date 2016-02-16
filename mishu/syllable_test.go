package mishu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyllables(t *testing.T) {
	//	assert.Equal(t, []string{}, Syllables(""));
	assert.Equal(t, []string{"pa", "ra", "le", "li", "pi", "ped"}, Syllables("paralelipiped"))
	assert.Equal(t, []string{"ma", "sa"}, Syllables("masa"))
	assert.Equal(t, []string{"pa", "du", "re"}, Syllables("padure"))
	assert.Equal(t, []string{"u", "ti", "li", "za", "re"}, Syllables("utilizare"))
	assert.Equal(t, []string{"pai", "ne"}, Syllables("paine"))
	assert.Equal(t, []string{"stro", "pea", "la"}, Syllables("stropeala"))
	assert.Equal(t, []string{"a", "xa"}, Syllables("axa"))
	assert.Equal(t, []string{"e", "xa", "men"}, Syllables("examen"))
	assert.Equal(t, []string{"ve", "che"}, Syllables("veche"))
	assert.Equal(t, []string{"in", "ger"}, Syllables("inger"))
	assert.Equal(t, []string{"un", "gher"}, Syllables("ungher"))
	assert.Equal(t, []string{"ar", "ta"}, Syllables("arta"))
	assert.Equal(t, []string{"ac", "tiv"}, Syllables("activ"))
	assert.Equal(t, []string{"mun", "te"}, Syllables("munte"))
	assert.Equal(t, []string{"un", "ghi", "e"}, Syllables("unghie"))
	assert.Equal(t, []string{"o", "braz"}, Syllables("obraz"))
	assert.Equal(t, []string{"co", "dru"}, Syllables("codru"))
	assert.Equal(t, []string{"a", "flu", "ent"}, Syllables("afluent"))
	assert.Equal(t, []string{"a", "gra", "fa"}, Syllables("agrafa"))
	assert.Equal(t, []string{"e", "vla", "vi", "e"}, Syllables("evlavie"))
	assert.Equal(t, []string{"as", "tru"}, Syllables("astru"))
	assert.Equal(t, []string{"mon", "stru"}, Syllables("monstru"))
	assert.Equal(t, []string{"sculp", "tu", "ra"}, Syllables("sculptura"))
	assert.Equal(t, []string{"somp", "tu", "os"}, Syllables("somptuos"))
	assert.Equal(t, []string{"linc", "sii"}, Syllables("lincsii"))
	assert.Equal(t, []string{"func", "ti", "e"}, Syllables("functie"))
	assert.Equal(t, []string{"ast", "ma", "tic"}, Syllables("astmatic"))
	assert.Equal(t, []string{"a", "er"}, Syllables("aer"))
	assert.Equal(t, []string{"a", "le", "e"}, Syllables("alee"))
	assert.Equal(t, []string{"po", "e", "zi", "e"}, Syllables("poezie"))
	assert.Equal(t, []string{"plo", "ua"}, Syllables("ploua"))
	assert.Equal(t, []string{"du", "ios"}, Syllables("duios"))
	assert.Equal(t, []string{"ploa", "ie"}, Syllables("ploaie"))
	assert.Equal(t, []string{"stea", "ua"}, Syllables("steaua"))
}
