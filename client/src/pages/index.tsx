import Link from "next/link";
import { useRouter } from "next/router";
import React, { ReactElement, useState } from "react";
import { useCookies } from "react-cookie";
import toast from "react-hot-toast";
import { useQuery } from "react-query";
import commafy from "../bin/commafy";
import Container from "../components/Layout/Container";
import PageContainer from "../components/Layout/PageContainer";
import UploadImage from "../components/UploadImage";
import { storeAPI } from "../services/api";

interface Props {}

function Home({}: Props): ReactElement {
  const { data } = useQuery("store", storeAPI.products.getList);
  if (!data) return <Container>Loading...</Container>;
  return (
    <Container>
      <PageContainer className="flex max-w-5xl flex-col min-h-[500px] px-3">
        <Products data={data} />
        <hr className="my-12" />
        <Form />
      </PageContainer>
    </Container>
  );
}

function Products({ data }: any) {
  return (
    <div className="mb-8">
      <h1 className="w-full mb-10 text-3xl font-bold text-center">Products 🎧</h1>
      {data.length === 0 && <div className="w-full p-3 text-yellow-800 bg-yellow-100 border border-yellow-400">No products yet, you can add products as long as you're logged in!</div>}
      {data.map((product: any) => (
        <Link key={product.product_slug} href={"/" + product.product_slug}>
          <a className="">
            <div className="flex flex-col justify-between p-4 mb-8 border hover:bg-slate-50 sm:flex-row rounded-2xl">
              <div>
                <div className="mb-2 text-5xl font-bold">{product.name}</div>
                <div className="mb-2 text-lg">{product.description}</div>
                <div className="mb-4 text-lg">IDR {commafy(product.price)}</div>
              </div>
              <img
                className="max-w-xs border shadow-xl rounded-xl"
                src={product.image_url}
                alt={data.description}
              />
            </div>
          </a>
        </Link>
      ))}
    </div>
  );
}
function Form() {
  const [imageURL, setImageURL] = useState("");

  const router = useRouter();
  const [form, setForm] = useState({
    name: "",
    description: "",
    product_slug: "",
    price: 0,
    image_url: "",
  });

  function handleChange(e: any) {
    const name = e.target.name;
    const value = e.target.value;
    setForm({ ...form, [name]: value });
  }

  const [cookies] = useCookies(["token"]);
  async function handleSubmit(e: any) {
    e.preventDefault();
    try {
      console.log(imageURL);
      const body = {
        image_url: imageURL,
        name: form.name,
        description: form.description,
        price: Number(form.price),
      };
      const res = await storeAPI.products.create(body, cookies.token);
      console.log(res);
      router.push(`/${res.product_slug}`);
    } catch (error) {
      console.error(error);
      toast.error("Error creating product, please try again later");
    }
  }
  const inputClassName = `my-2 w-full border px-3 py-2`;
  return (
    <div>
      <h1 className="mb-3 text-3xl font-bold">Submit Product 👇</h1>
      <form onSubmit={handleSubmit}>
        <input
          type={"text"}
          name="name"
          placeholder="Product Name"
          className={inputClassName}
          onChange={handleChange}
        />
        <input
          type={"text"}
          name="description"
          placeholder="Description"
          className={inputClassName}
          onChange={handleChange}
        />
        <input
          type={"text"}
          name="price"
          placeholder="Price (IDR)"
          className={inputClassName}
          onChange={handleChange}
        />
        <UploadImage imageURL={imageURL} setImageURL={setImageURL} />
        <div className="flex justify-center">
          <button
            type="submit"
            className="w-full px-4 py-3 mt-3 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
}

export default Home;
