import { useRouter } from "next/router";
import { useQuery } from "react-query";
import Container from "../components/Layout/Container";
import { storeAPI } from "../services/api";
import toast from "react-hot-toast";
import commafy from "../bin/commafy";
import {useCookies} from "react-cookie";
export default function ProductSlugPage() {
  const router = useRouter();
  const product_slug = router.query.product_slug as string;
const [cookies] = useCookies(['token']);
  const { data } = useQuery(["product", router], async () =>
    storeAPI.products.slug(product_slug)
  );
  async function handleBuy(e: any) {
    e.preventDefault();
    const req = await storeAPI.purchase(product_slug, cookies.token);
    console.log(req);
    toast.success(`You bought ${data.name} for ${commafy(data.price)}`);
    router.push("/");
  }

  if (!data)
    return (
      <Container>
        <h1>Loading...</h1>
      </Container>
    );

  return (
    <Container>
      <div>
        <div className="flex justify-between max-w-4xl mx-auto">
          <div>
            <div className="mb-2 text-5xl">{data.name}</div>
            <div className="mb-2 text-lg">{data.description}</div>
            <div className="mb-4 text-lg">IDR {commafy(data.price)}</div>
            <button
              className="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
              onClick={handleBuy}
            >
              Buy Now
            </button>
          </div>
          <img
            className="max-w-xs border shadow-xl rounded-xl"
            src={data.image_url}
            alt={data.description}
          />
        </div>
      </div>
    </Container>
  );
}
