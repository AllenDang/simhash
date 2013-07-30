simhash
=======

A library to find the percentage of similarity between two given strings (can be expanded to compare every thing!).

This is a port of C# version SimHash created by ArefKarimi, you can find it here (http://simhash.codeplex.com/).

Usage
-----
    needle := "Reading bytes into structs using reflection"
    hayStack := "Golang - mapping an variable length array to a struct"

    likeness := GetLikenessValue(needle, hayStack)
    fmt.Println("Likeness:", likeness)
