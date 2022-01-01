package ngram_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/rk-ee/virkvirivirvavirn/pkg/ngram"
	"github.com/stretchr/testify/assert"
)

func TestDecodeNgram(t *testing.T) {
	t.Parallel()

	input := strings.NewReader(
		"EuroCom	1878,1,1	1967,1,1	1968,1,1	1972,9,1	1977,3,2	1979,8,5	1980,2,2	1982,1,1	1986,10,7	1987,1,1	1988,2,2	1989,13,3	1993,2,2	1994,1,1	1995,29,7	1996,7,3	1997,6,5	1998,12,12	1999,45,13	2000,18,7	2001,83,19	2002,47,14	2003,230,13	2004,40,20	2005,10,6	2006,37,15	2007,154,20	2008,151,15	2009,78,21	2010,63,31	2011,55,27	2012,90,14	2013,56,17	2014,89,11	2015,13,6	2016,125,9	2017,4,3	2018,45,6	2019,18,4\n" +
			"GBHCs_NOUN	1988,24,3	1991,23,2	1993,27,4	1994,18,18	1997,3,3	1998,8,3	1999,84,4	2000,13,5	2001,18,8	2002,12,4	2003,12,4	2004,10,2	2005,3,1	2006,18,6	2007,20,7	2012,2,1	2013,4,1\n" +
			"Georgsmarien_NOUN	1901,4,2	1906,2,2	1907,1,1	1908,5,5	1909,1,1	1910,1,1	1912,1,1	1914,1,1	1919,11,11	1924,2,1	1934,2,2	1935,2,2	1947,1,1	1951,1,1	1954,1,1	1955,1,1	1960,3,3	1962,1,1	1964,3,3	1966,2,1	1971,2,2	1974,1,1	1979,3,3	1980,2,2	1981,1,1	1987,2,2	1992,1,1	1995,1,1	1996,1,1	2001,3,3	2003,1,1	2004,1,1	2012,1,1	2014,2,2\n",
	)
	expected := []ngram.Ngram{
		{Value: "EuroCom", Type: "", Years: map[int]ngram.NYear{1878: {1, 1}, 1967: {1, 1}, 1968: {1, 1}, 1972: {9, 1}, 1977: {3, 2}, 1979: {8, 5}, 1980: {2, 2}, 1982: {1, 1}, 1986: {10, 7}, 1987: {1, 1}, 1988: {2, 2}, 1989: {13, 3}, 1993: {2, 2}, 1994: {1, 1}, 1995: {29, 7}, 1996: {7, 3}, 1997: {6, 5}, 1998: {12, 12}, 1999: {45, 13}, 2000: {18, 7}, 2001: {83, 19}, 2002: {47, 14}, 2003: {230, 13}, 2004: {40, 20}, 2005: {10, 6}, 2006: {37, 15}, 2007: {154, 20}, 2008: {151, 15}, 2009: {78, 21}, 2010: {63, 31}, 2011: {55, 27}, 2012: {90, 14}, 2013: {56, 17}, 2014: {89, 11}, 2015: {13, 6}, 2016: {125, 9}, 2017: {4, 3}, 2018: {45, 6}, 2019: {18, 4}}},
		{Value: "GBHCs", Type: "NOUN", Years: map[int]ngram.NYear{1988: {24, 3}, 1991: {23, 2}, 1993: {27, 4}, 1994: {18, 18}, 1997: {3, 3}, 1998: {8, 3}, 1999: {84, 4}, 2000: {13, 5}, 2001: {18, 8}, 2002: {12, 4}, 2003: {12, 4}, 2004: {10, 2}, 2005: {3, 1}, 2006: {18, 6}, 2007: {20, 7}, 2012: {2, 1}, 2013: {4, 1}}},
		{Value: "Georgsmarien", Type: "NOUN", Years: map[int]ngram.NYear{1901: {4, 2}, 1906: {2, 2}, 1907: {1, 1}, 1908: {5, 5}, 1909: {1, 1}, 1910: {1, 1}, 1912: {1, 1}, 1914: {1, 1}, 1919: {11, 11}, 1924: {2, 1}, 1934: {2, 2}, 1935: {2, 2}, 1947: {1, 1}, 1951: {1, 1}, 1954: {1, 1}, 1955: {1, 1}, 1960: {3, 3}, 1962: {1, 1}, 1964: {3, 3}, 1966: {2, 1}, 1971: {2, 2}, 1974: {1, 1}, 1979: {3, 3}, 1980: {2, 2}, 1981: {1, 1}, 1987: {2, 2}, 1992: {1, 1}, 1995: {1, 1}, 1996: {1, 1}, 2001: {3, 3}, 2003: {1, 1}, 2004: {1, 1}, 2012: {1, 1}, 2014: {2, 2}}},
	}

	got, err := ngram.DecodeNgram(input)
	assert.Nil(t, err)
	fmt.Println(len(got))
	sort.SliceStable(got, func(i, j int) bool { return got[i].Value < got[j].Value })
	assert.Equal(t, expected, got)
}

func TestErrDecode(t *testing.T) {
	t.Parallel()

	input := strings.NewReader("xyz")
	_, err := ngram.DecodeNgram(input)
	assert.ErrorIs(t, err, ngram.ErrFormat)
}