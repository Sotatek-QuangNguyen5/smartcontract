solc:
	solc --abi --bin --overwrite contract/contract.sol -o build
	
abigen:
	abigen --bin=build/MyContract.bin --abi=build/MyContract.abi --pkg=MyContract --out=mycontract/contract.go

deploy:
	go run deploy/main.go