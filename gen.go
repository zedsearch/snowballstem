package snowballstem

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/arabic.sbl -go -o arabic/arabic_stemmer -gop arabic -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w arabic/arabic_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/basque.sbl -go -o basque/basque_stemmer -gop basque -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w basque/basque_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/catalan.sbl -go -o catalan/catalan_stemmer -gop catalan -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w catalan/catalan_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/danish.sbl -go -o danish/danish_stemmer -gop danish -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w danish/danish_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/dutch.sbl -go -o dutch/dutch_stemmer -gop dutch -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w dutch/dutch_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/english.sbl -go -o english/english_stemmer -gop english -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w english/english_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/finnish.sbl -go -o finnish/finnish_stemmer -gop finnish -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w finnish/finnish_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/french.sbl -go -o french/french_stemmer -gop french -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w french/french_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/german.sbl -go -o german/german_stemmer -gop german -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w german/german_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/greek.sbl -go -o greek/greek_stemmer -gop greek -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w greek/greek_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/hindi.sbl -go -o hindi/hindi_stemmer -gop hindi -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w hindi/hindi_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/hungarian.sbl -go -o hungarian/hungarian_stemmer -gop hungarian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w hungarian/hungarian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/indonesian.sbl -go -o indonesian/indonesian_stemmer -gop indonesian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w indonesian/indonesian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/irish.sbl -go -o irish/irish_stemmer -gop irish -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w irish/irish_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/italian.sbl -go -o italian/italian_stemmer -gop italian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w italian/italian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/lithuanian.sbl -go -o lithuanian/lithuanian_stemmer -gop lithuanian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w lithuanian/lithuanian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/nepali.sbl -go -o nepali/nepali_stemmer -gop nepali -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w nepali/nepali_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/norwegian.sbl -go -o norwegian/norwegian_stemmer -gop norwegian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w norwegian/norwegian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/porter.sbl -go -o porter/porter_stemmer -gop porter -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w porter/porter_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/portuguese.sbl -go -o portuguese/portuguese_stemmer -gop portuguese -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w portuguese/portuguese_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/romanian.sbl -go -o romanian/romanian_stemmer -gop romanian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w romanian/romanian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/russian.sbl -go -o russian/russian_stemmer -gop russian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w russian/russian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/serbian.sbl -go -o serbian/serbian_stemmer -gop serbian -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w serbian/serbian_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/spanish.sbl -go -o spanish/spanish_stemmer -gop spanish -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w spanish/spanish_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/swedish.sbl -go -o swedish/swedish_stemmer -gop swedish -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w swedish/swedish_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/tamil.sbl -go -o tamil/tamil_stemmer -gop tamil -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w tamil/tamil_stemmer.go

//go:generate $SNOWBALL/snowball $SNOWBALL/algorithms/turkish.sbl -go -o turkish/turkish_stemmer -gop turkish -gor github.com/blevesearch/snowballstem
//go:generate gofmt -s -w turkish/turkish_stemmer.go

