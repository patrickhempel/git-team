#!/usr/bin/env bats

load '/bats-libs/bats-support/load.bash'
load '/bats-libs/bats-assert/load.bash'

setup() {
	/usr/local/bin/git-team add a 'A <a@x.y>'
	/usr/local/bin/git-team add b 'B <b@x.y>'
	/usr/local/bin/git-team add c 'C <c@x.y>'
}

@test "git-team: assignments (default) should show all alias -> coauthor assignments" {
  run /usr/local/bin/git-team assignments
  assert_success
  assert_line --index 0 'Assignments:'
  assert_line --index 1 '------------'
  assert_line --index 2 "'a' -> 'A <a@x.y>'"
  assert_line --index 3 "'b' -> 'B <b@x.y>'"
  assert_line --index 4 "'c' -> 'C <c@x.y>'"
}

@test "git-team: assignments ls should show all alias -> coauthor assignments" {
  run /usr/local/bin/git-team assignments ls
  assert_success
  assert_line --index 0 'Assignments:'
  assert_line --index 1 '------------'
  assert_line --index 2 "'a' -> 'A <a@x.y>'"
  assert_line --index 3 "'b' -> 'B <b@x.y>'"
  assert_line --index 4 "'c' -> 'C <c@x.y>'"
}

@test "git-team: assignments list should show all alias -> coauthor assignments" {
  run /usr/local/bin/git-team assignments ls
  assert_success
  assert_line --index 0 'Assignments:'
  assert_line --index 1 '------------'
  assert_line --index 2 "'a' -> 'A <a@x.y>'"
  assert_line --index 3 "'b' -> 'B <b@x.y>'"
  assert_line --index 4 "'c' -> 'C <c@x.y>'"
}

teardown() {
	/usr/local/bin/git-team rm a
	/usr/local/bin/git-team rm b
	/usr/local/bin/git-team rm c
}
