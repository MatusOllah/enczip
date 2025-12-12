# 📦 enczip

**enczip** is a fork of the Go stdlib package `archive/zip` to add support for various text encodings other than just UTF-8.

## Why?

> ⚠️ rant incoming

Because Shift-JIS ZIP archives exist.
Sometimes, especially if you tinker with Japanese singing voice synthesis, life hands you an ancient ZIP archive that has been faithfully preserved in pure, glorious Shift-JIS. And sometimes that ZIP archive belongs to a certain UTAU voicebank (*wink wink, Yamine Renri*).

The Go stdlib's `archive/zip` package has a strict "UTF-8 or else" policy, which is great in theory, until you're having to read the most Shift-JIS'd, mojibaked, Go-unfriendly ZIP archive ever produced by 2000s-era software that apparently never got the memo. This happened to me while working on [one of my other projects](https://github.com/MatusOllah/gotau).
The result? Mojibake, random errors, confusion, and a lot of time spent questioning why I chose this hobby and why I didn't just cough up the ~100 EUR for a more commercial SVS engine like Synthesizer V and call it a day.

So, I forked `archive/zip`, lifted the UTF-8-only restriction, and added better support for obscure text encodings beyond just UTF-8.

Thus, **enczip** - a fork of `archive/zip` with proper handling for non-UTF-8 filename encodings.

tl;dr: Shift-JIS sucks. UTAU should use UTF-8 already.
