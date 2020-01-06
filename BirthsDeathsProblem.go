package main

//Given list of births and deaths, determine years where most people were alive
//https://stackoverflow.com/questions/32896439/given-list-of-births-and-deaths-determine-years-where-most-people-were-alive
//https://www.presidentsusa.net/birth.html
import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name  string
		Born  int
		Death int
	}{
		//Expected results are 1919 and 1920.
		{"Nicolas", 1900, 1975},
		{"Vladimir", 1970, 2000},
		{"Julius", 1950, 1985},
		{"Alexander", 1900, 1920},
		{"Obama", 1910, 1920},
		{"George", 1915, 1920},
		{"Benjamin", 1919, 1925},
		{"Donald", 1918, 1918},
	}
	/* Results:
		years [{0 0} {1929 1} {1933 1} {1941 1} {1990 1} {1992 1} {1996 1} {1928 1} {1942 1} {1985 1} {1931 1} {1938 1} {1998 1}
		{1940 1} {1999 1} {1945 1} {1946 1} {1994 1} {1934 1} {1944 1} {1987 1} {1995 1} {1943 1} {1948 1} {1926 1} {1935 1} 
		{1997 1} {1947 1} {1993 1} {1930 1} {1939 1} {1988 1} {1925 1} {1927 1} {1989 1} {1991 1} {1932 1} {1949 1} {1936 1} 
		{1937 1} {1986 1} {1905 2} {1976 2} {1982 2} {1904 2} {1924 2} {1959 2} {1956 2} {1960 2} {1980 2} {1906 2} {1968 2} 
		{1907 2} {1964 2} {1967 2} {1977 2} {1969 2} {1920 2} {1921 2} {1965 2} {1902 2} {1909 2} {1922 2} {1961 2} {1900 2} 
		{1901 2} {1983 2} {1954 2} {1903 2} {1981 2} {1953 2} {1962 2} {1978 2} {1952 2} {1963 2} {1966 2} {1923 2} {1951 2} 
		{1957 2} {1975 2} {1908 2} {1984 2} {1950 2} {1955 2} {1958 2} {1979 2} {1911 3} {1970 3} {1972 3} {1974 3} {1912 3} 
		{1971 3} {1910 3} {1913 3} {1973 3} {1914 3} {1916 4} {1915 4} {1918 4} {1917 4} {1919 5}]

	*/

	t := make(map[int]int)

	//O(people*80)
	for _, v := range people {
		for i := v.Born; i < v.Death; i++ {
			t[i]++
		}
	}
	idx := make([]struct {
		Year  int
		Count int
	}, 1)
	// y = unique years in people
	// O(y)
	for k := range t {
		n := idx[0]
		n.Year = k
		n.Count = t[k]
		idx = append(idx, n)
	}
	//sort by count
	sort.SliceStable(idx, func(i, j int) bool {
		return idx[i].Count < idx[j].Count
	})

	fmt.Printf("years %v\n", idx)

}
