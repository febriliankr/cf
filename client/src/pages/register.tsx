
import { useRouter } from "next/router";
import { ReactElement, useState } from "react";
import toast from "react-hot-toast";
import Container from "../components/Layout/Container";
import PageContainer from "../components/Layout/PageContainer";
import { storeAPI } from "../services/api";

interface Props {}

export default function RegisterPage({}: Props): ReactElement {
  const [form, setForm] = useState({
    student_id: "",
    password: "",
  });

  const router = useRouter();
  async function handleSubmit(e: any) {
    e.preventDefault();
    try {
      const req = await storeAPI.user.register(form.student_id, form.password);
      document.cookie="token="+req.token;
      console.log(req);
      router.push("/");
    } catch (error: any) {
      console.error(error)
      toast.error(error.response?.data?.message);
      
    }
  }

  const inputClassName = `my-2 w-full border px-3 py-2`;
  return (
    <Container>
      <PageContainer className="flex items-center flex-col min-h-[500px] px-3">
        <div>
          <h1 className="mb-4 text-3xl font-bold text-center">Register</h1>
          <form onSubmit={handleSubmit}>
            <label htmlFor="student_id">Student ID</label>
            <input
              autoComplete="off"
              type="student_id"
              id="student_id"
              name="student_id"
              className={inputClassName}
              onChange={(e) =>
                setForm((prev) => ({
                  ...prev,
                  [e.target.name]: e.target.value,
                }))
              }
            />
            <label htmlFor="password">Password</label>
            <input
              type="password"
              autoComplete="off"
              className={inputClassName}
              id="password"
              name="password"
              onChange={(e) =>
                setForm((prev) => ({
                  ...prev,
                  [e.target.name]: e.target.value,
                }))
              }
            />
            <button
              type="submit"
              onClick={handleSubmit}
              className="w-full px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
            >
              Register
            </button>
          </form>
        </div>
      </PageContainer>
    </Container>
  );
}
