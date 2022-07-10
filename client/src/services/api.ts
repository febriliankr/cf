import Axios from "axios";

const api = Axios.create({
  baseURL: process.env.API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

async function login(student_id: string, password: string) {
  const body = {
    student_id: Number(student_id),
    password,
  };
  const { data } = await api.post("/user/login", body);
  return data;
}

async function register(student_id: string, password: string) {
  const body = {
    student_id: Number(student_id),
    password,
  };
  const { data } = await api.post("/user", body);
  return data;
}

async function getList() {
  const { data } = await api.get("/product");
  return data.products || [];
}

type ProductRequest = {
  name: string;
  description: string;
  price: number;
  image_url: string;
};

async function create(body: ProductRequest, token: string) {
  const { data } = await api.post(
    "/product",
    body,

    { headers: { Authorization: "Bearer " + token } }
  );
  return data;
}

async function slug(slug: string) {
  const { data } = await api.get(`/product/${slug}`);
  return data;
}

async function upload(file: any) {
  const formData = new FormData();
  formData.append("file", file);
  formData.append("filename", file.name);
  const url = `/product/file`;
  const { data } = await api.post(url, formData);
  return data;
}

async function purchase(product_slug: string, token: string) {
  const { data } = await api.post(
    `/purchase`,
    { product_slug: product_slug },
    { headers: { Authorization: "Bearer " + token } }
  );
  return data;
}

const balance = {
  get: getBalance,
  add: addBalance,
  withdraw: withdrawBalance,
};

async function getBalance() {
  const { data } = await api.get(`/balance`);
  return data.balance;
}

async function addBalance(amount: number, token: string) {
  const { data } = await api.patch(
    `/balance`,
    {
      amount: amount,
      type: "deposit",
    },

    { headers: { Authorization: "Bearer " + token } }
  );
  return data.balance;
}
async function withdrawBalance(amount: number, token: string) {
  const { data } = await api.patch(
    `/balance`,
    {
      amount: amount,
      type: "withdraw",
    },

    { headers: { Authorization: "Bearer " + token } }
  );
  return data.balance;
}
export const storeAPI = {
  user: {
    login,
    register,
  },
  products: {
    getList,
    slug,
    create,
    upload,
  },
  purchase,
  balance: balance,
};
