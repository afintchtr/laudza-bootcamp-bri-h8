const printData = (payload) => {
  console.log(payload);
};

const fetchData = async (param, callback) => {
  try {
    const response = await fetch(`https://api.github.com/users/${param}`);
    const data = await response.json();
    callback(data);
  } catch (error) {
    console.log(error);
  }
};

const fetchUserData = (username, fetchHandler) => {
  fetchHandler(username, printData);
};

fetchUserData("afintchtr", fetchData);
