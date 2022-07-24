package main

func Example_main() {
	goMain([]string{"mywc", "testdata/wc/humpty_dumpty.txt"})
	// Output:
	//       lines      words characters      bytes
	//           4         26        142        142 testdata/wc/humpty_dumpty.txt
}
