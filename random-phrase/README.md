# random-phrase

This is a helper for generating random passphrases. It's inspired by the famous
XKCD "correct hosrse battery staple" comic, <https://xkcd.com/936/>.

The default word list is /usr/share/dict/words (which is pruned a bit). There
is also a helper to download the EFF's "long" diceware wordlist.

Based on some simple calculations (which you should regard with suspicion; this
is not my day job), we also print some stats about the entropy contained in the
generated passphrases, as well as a nice comparison to equivalent random
passwords.
