go build -o build/
for i in {1..13}
do
echo "TEST: "  $i
build/progetto < tests/"input-"$i".txt" > "output.txt"
diff "output.txt" tests/"expected-"$i".txt"
rm "output.txt"
echo "----------"
done