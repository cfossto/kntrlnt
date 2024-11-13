# kntrlnt
KNTRLNT (Kontrollant) is a web file browser written in Go.


# How to build this piece of magic
```docker build -t kntrlnt .```

# How to run this marvelous thing
```docker run -p 8001:8001 -v <"insert_your_base_folder">:/data/kntrlnt kntrlnt```

It is then served on port 8001. That can be changed, but by default it is run internally on 8001.
