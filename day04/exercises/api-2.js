const printData = (payload) => {
  console.log(payload);
};

const fetchData = async (param, callback) => {
  return new Promise((resolve, reject) => {
    fetch(`https://api.github.com/users/${param}`)
      .then((response) => {
        if (!response.ok) console.log(`HTTP error! Status: ${response.status}`);
        return response.json();
      })
      .then((data) => resolve(callback(data)))
      .catch((error) => reject(error));
  });
};

const fetchUserData = (username, fetchHandler) => {
  fetchHandler(username, printData);
};

fetchUserData("afintchtr", fetchData);
