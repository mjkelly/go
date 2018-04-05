#!/bin/bash
# Downloads EFF's "long" wordlist (meant for literal diceware passwords, but
# also useful here).
#
# This can be an alternative to /usr/share/dict/words.

f=eff_long_list.txt
curl --silent --show-error \
  https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt \
  | awk '{print $2}' > $f
echo wrote $f
