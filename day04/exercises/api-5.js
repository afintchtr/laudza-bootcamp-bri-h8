const fetchSomething = async (param) => {
  try {
    const response = await fetch(param);
    const data = await response.json();
    console.log(data.slice(0, 2));
  } catch (err) {
    console.error(`Error in ${param}: ${err}`);
  }
};

const postsUrl = "https://jsonplaceholder.typicode.com/posts";
const commentsUrl = "https://jsonplaceholder.typicode.com/comments";

const fetchAll = async (fetchHandler) => {
  await fetchHandler(postsUrl);
  await fetchHandler(commentsUrl);
  console.log("\nDone!");
};

fetchAll(fetchSomething);
