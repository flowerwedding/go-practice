package main

import "fmt"

func main(){
	beginWord:="green"
	wordList := []string{"hit","tag","green","date","early","need","yet"}
	jielong(beginWord,wordList)
}

func jielong(beginWord string,wordList []string){
    ks:=beginWord[len(beginWord)-1:]
    js:=beginWord[0:1]
    var ans []string =make([]string,1,1)
    ans[0] = beginWord
    for {
		if js != ks {
			for i := 0; i < len(wordList); i++ {
				next := wordList[i]
				jl := next[0:1]
				if jl == ks {
					ks = next[len(next)-1:]
					ans = append(ans, next)
					break
				}
			}
		} else {
			fmt.Println(ans)
			break
		}
	}
}