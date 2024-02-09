go build
for i in {1..7}
do
echo "TEST: "  $i
build/progetto < tests/"input-"$i".txt" > "output.txt"
diff -a "output.txt" tests/"expected-"$i".txt"
rm "output.txt"
echo "----------"
done