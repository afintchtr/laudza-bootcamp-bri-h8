findFrequency = (someArray) => {
  let object = {};
  someArray.forEach((element) => {
    if (typeof element === "string") element = element + " but string";
    if (element.toString() in object) object[element.toString()]++;
    else {
      object[element.toString()] = 1;
    }
  });
  maxValue = 0;
  maxIndex = null;
  Object.entries(object).forEach(([key, value]) => {
    if (value > maxValue) {
      maxIndex = key;
      maxValue = value;
    }
  });
  console.log(`highest number is ${maxIndex} with frequency: ${maxValue}`);
};

findFrequencyButWithReduce = (someArray) => {
  let resultObject = someArray.reduce((object, key) => {
    key = typeof key === "string" ? key + " but string" : key.toString();
    if (!object[key]) object[key] = 1;
    else object[key]++;
    return object;
  }, {});
  // console.log(resultObject);
  maxValue = 0;
  maxIndex = null;
  Object.entries(resultObject).forEach(([key, value]) => {
    if (value > maxValue) {
      maxIndex = key;
      maxValue = value;
    }
  });
  console.log(`highest number is ${maxIndex} with frequency: ${maxValue}`);
};

input1 = [1, 2, 2, 3, 4, 2];
input2 = [5, 5, 7, "a", "b", "a", "a", "a", "c", 7, 5, 5, "5"];

findFrequency(input1);
findFrequency(input2);
findFrequencyButWithReduce(input1);
findFrequencyButWithReduce(input2);
