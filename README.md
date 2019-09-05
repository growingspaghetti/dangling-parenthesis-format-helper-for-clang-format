# dangling-parenthesis-format-helper-for-clang-format

"add option for dangling parenthesis" https://reviews.llvm.org/D33029

```
cat a.java | gangle

or 

dangle a.java
```

from

```
    static Optional<String> getPreviousDC(
        int idInt) {
        ...
    }

    static Optional<String> getDaPreviousDC(
        int idInt) throws IllegaldadadaException {
        ...
    }

    private static void fetchScreenShot(
        String url,
        String filePath) throws MalformedURLException,
                                IllegalStateException {
        ...
    }
```

to

```
    static Optional<String> getPreviousDC(
        int idInt
    ) {
        ...
    }

    static Optional<String> getDaPreviousDC(
        int idInt
    ) throws IllegaldadadaException {
        ...
    }

    private static void fetchScreenShot(
        String url,
        String filePath
    ) throws MalformedURLException,
             IllegalStateException
    {
        ...
    }
```

