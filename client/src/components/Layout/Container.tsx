import React, { ReactElement, useState } from "react";
import Head from "next/head";
import Footer from "./Footer";
import Navbar from "./Navbar";
import Seo from "../SEO";
import classNames from "classnames";

import { useCookies } from "react-cookie";
import { useQuery } from "react-query";
import jwtDecode from "jwt-decode";
import { useRouter } from "next/router";
import { storeAPI } from "../../services/api";

interface Props {
  title?: string;
  children: ReactElement | string;
  className?: string;
}

function Container({ title, children, className = "" }: Props): ReactElement {
  const [cookies] = useCookies(["token"]);

  const [loadingButton, setLoadingButton] = useState({
    add: false,
    withdraw: false,
  });

  const [amount, setAmount] = useState(5000);
  async function handleAddBalance(e: any) {
    e.preventDefault();
    if (loadingButton.add) return;
    try {
      setLoadingButton({ ...loadingButton, add: true });
      const req = await storeAPI.balance.add(amount, cookies.token);
      console.log(req);
      refetch();
      setLoadingButton({ ...loadingButton, add: false });
    } catch (error: any) {
      console.error(error);
    }
  }

  async function handleWithdrawBalance(e: any) {
    e.preventDefault();
    if (loadingButton.withdraw) return;
    try {
      setLoadingButton({ ...loadingButton, withdraw: true });
      const req = await storeAPI.balance.withdraw(amount, cookies.token);
      console.log(req);
      refetch();
      setLoadingButton({ ...loadingButton, withdraw: false });
    } catch (error: any) {
      console.error(error);
    }
  }

  const { data: balance, refetch } = useQuery("balance", async () =>
    storeAPI.balance.get()
  );

  const { data: student_id } = useQuery(["profile", cookies], async () => {
    if (cookies.token) {
      const user: any = jwtDecode(cookies.token);
      if (user.exp) {
        const now = new Date().getTime();
        const isExpired = user.exp * 1000 - now < 0;
        const expiryEmoji = isExpired ? "🙏" : "🔥";
        const hoursToExpiry = (user.exp * 1000 - now) / (1000 * 3600);
        console.log(expiryEmoji, hoursToExpiry);
        return user.student_id;
      } else {
        // What if the user does not have an expiry key in their JWT?
        // IDK! All the newest one should have on
      }

      return user;
    }
  });
  return (
    <>
      <Seo title={title} />
      <Navbar
        loadingButton={loadingButton}
        student_id={student_id}
        balance={balance}
        handleAddBalance={handleAddBalance}
        handleWithdrawBalance={handleWithdrawBalance}
        amount={amount}
        setAmount={setAmount}
      />
      <div className={classNames("mt-64 sm:mt-56 px-3", className)}>
        {children}
      </div>
      <Footer student_id={student_id} />
    </>
  );
}

export default Container;
