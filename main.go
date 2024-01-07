package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/paulmach/orb/geojson"
)

func main() {
	flag.Parse()
	f, err := readStreamToFeatureCollection(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	j, err := json.Marshal(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(j))
}

func readStreamToFeatureCollection(reader io.Reader) (*geojson.FeatureCollection, error) {
	breader := bufio.NewReader(reader)

	fc := geojson.NewFeatureCollection()

	for {
		read, err := breader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, os.ErrClosed) || errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		feat, err := geojson.UnmarshalFeature(read)
		if err != nil {
			return nil, err
		}

		fc.Append(feat)
	}

	return fc, nil
}
