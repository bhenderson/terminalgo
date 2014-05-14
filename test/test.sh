set -e
cmd="test/test"

read -r -d '' EXP <<-'EOE' || true
stdin is a terminal
stdout is a terminal
-->break<--
stdout is a terminal
-->break<--
stdin is a terminal
EOE

run() {
  script -c "$1" -q /dev/null
}

test_out() {
  read -d '' ACT || true
  diff <(echo "$EXP") <(echo "$ACT")
}

lb="echo '-->break<--'"

run "$cmd; $lb; echo | $cmd; $lb; $cmd | head" | test_out
