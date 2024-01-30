// const outputOnModal = () => {
//   $("#outputModal").modal("show");
// };

// outputOnModal();

const operator = document.getElementsByClassName("btn btn-primary");
const operation = (oper1, oper2, operat) => {
  switch (operat) {
    case 0:
      return oper1 + oper2;
    case 1:
      return oper1 - oper2;
    case 2:
      return oper1 * oper2;
    case 3:
      return oper1 / oper2;
  }
};

for (let i = 0; i < 4; i++) {
  operator[i].addEventListener("click", () => {
    let firstNumber = document.getElementById("angka1");
    let firstNumberValue = firstNumber.value;

    let secondNumber = document.getElementById("angka2");
    let secondNumberValue = secondNumber.value;

    if (firstNumberValue == "" || secondNumberValue == "") {
      console.log("Pls fill the field correctly");
      return;
    }

    let result = operation(
      parseInt(firstNumberValue),
      parseInt(secondNumberValue),
      i
    );
    console.log(result);
  });
}

const check = document.getElementById("checkCreds");
check.addEventListener("click", () => {
  let uname = document.getElementById("username").value;
  let pass = document.getElementById("password").value;

  if (uname == "admin" && pass == "123") {
    window.open("http://stackoverflow.com", "_blank");
  } else alert("Wrong Credentials!");
});

const voila = document.getElementById("voila1");
voila.addEventListener("click", () => {
  let row = document.getElementById("stars").value;
  if (row == "") console.log("jangan lupa isi bintang!");
  for (let i = 1; i <= row; i++) {
    temp = "";
    for (let j = row; j > 0; j--) {
      if (j <= i) temp += "*";
      else temp += " ";
    }
    console.log(temp);
  }
});

const voilaButWithHashtag = document.getElementById("voila2");
voilaButWithHashtag.addEventListener("click", () => {
  let row = document.getElementById("stars").value;
  if (row == "") console.log("jangan lupa isi bintang!");
  for (let i = 1; i <= row; i++) {
    temp = "";
    for (let j = row; j > 0; j--) {
      if (j <= i) {
        if (row % 2 == 0) {
          if (j % 2 == 0) temp += "#";
          else temp += "*";
        } else {
          if (j % 2 == 0) temp += "#";
          else temp += "*";
        }
      } else temp += " ";
    }
    console.log(temp);
  }
});

const voilaButEquillateral = document.getElementById("voila3");
voilaButEquillateral.addEventListener("click", () => {
  let row = document.getElementById("stars").value;
  if (row == "") console.log("jangan lupa isi bintang!");
  for (let i = 1; i <= row; i++) {
    temp = "";
    for (let j = 1; j <= row - i; j++) temp += " ";
    for (let j = 1; j <= i * 2 - 1; j++) temp += "*";
    console.log(temp);
  }
});

const sortingButtonUp = document.getElementById("sortingButtonUp");
sortingButtonUp.addEventListener("click", () => {
  let input = document.getElementById("unsorted-numbers").value;
  if (input == "") console.log("Isi dlu angkanya!");
  let someArray = input.split(" ");
  someArray = someArray.map((item) => parseInt(item, 10));
  for (let i = 0; i < someArray.length; i++) {
    for (let j = 0; j < someArray.length - 1; j++) {
      if (someArray[j] > someArray[j + 1]) {
        temp = someArray[j];
        someArray[j] = someArray[j + 1];
        someArray[j + 1] = temp;
      }
    }
  }
  console.log(someArray);
});

const sortingButtonDown = document.getElementById("sortingButtonDown");
sortingButtonDown.addEventListener("click", () => {
  let input = document.getElementById("unsorted-numbers").value;
  if (input == "") console.log("Isi dlu angkanya!");
  let someArray = input.split(" ");
  someArray = someArray.map((item) => parseInt(item, 10));
  for (let i = 0; i < someArray.length; i++) {
    for (let j = 0; j < someArray.length - 1; j++) {
      if (someArray[j] < someArray[j + 1]) {
        temp = someArray[j];
        someArray[j] = someArray[j + 1];
        someArray[j + 1] = temp;
      }
    }
  }
  console.log(someArray);
});

const reverseButton = document.getElementById("reverseButton");
reverseButton.addEventListener("click", () => {
  let input = document.getElementById("reversable-text").value;
  if (input == "") console.log("Isi dlu string-nya!");
  let someArray = input.split("");
  let isPalindrom = true;
  for (let i = 0; i < Math.floor(someArray.length / 2); i++) {
    if (someArray[i] != someArray[someArray.length - 1 - i])
      isPalindrom = false;
    let temp = someArray[i];
    someArray[i] = someArray[someArray.length - 1 - i];
    someArray[someArray.length - 1 - i] = temp;
  }
  console.log(someArray.join(""));
  if (isPalindrom == true && input != "")
    console.log("that string is a Palindrom!");
});

const findBigThree = document.getElementById("findBigThree");
findBigThree.addEventListener("click", () => {
  let input = document.getElementById("some-array").value;
  let numArray = input.split(" ");
  numArray = numArray.map((item) => parseInt(item, 10));
  if (numArray.length < 3) {
    console.log("Error! Array-nya kurang");
  } else {
    let numSet = [...new Set(numArray)];
    let lengthOfSet = numSet.length;
    numSet.sort((a, b) => b - a);
    let result = 0;
    if (lengthOfSet == 1) result = numSet[0];
    else if (lengthOfSet == 2) result = numSet[0] * numSet[1];
    else result = numSet[0] * numSet[1] * numSet[2];
    console.log(result);
  }
});
