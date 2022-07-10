import React, { ReactElement, useState } from "react";
import Link from "next/link";
import BurgerMenu from "../BurgerMenu";
import links from "./data/links";
import { useRouter } from "next/router";
import commafy from "../../bin/commafy";

function Navbar({
  student_id,
  balance,
  handleAddBalance,
  handleWithdrawBalance,
  loadingButton,
  amount,
  setAmount,
}: any): ReactElement {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const router = useRouter();
  function handleChangeAmount(e: any) {
    const value = Number(e.target.value);
    if (value > 0) {
      setAmount(value);
    } else {
      setAmount(0);
    }
  }
  return (
    <nav className="absolute top-0 z-10 w-full border-b">
      <ul className="flex flex-col items-center justify-between max-w-6xl px-8 py-6 mx-auto sm:flex-row">
        <li className="">
          <Link href="/">
            <a className="">
              <h1 className="text-2xl font-extrabold">Kantin Kejujuran 🦋</h1>
              <h2 className="mb-2 text-lg font-semibold">
                Balance 💸 IDR {commafy(balance)}
              </h2>
              <div className="px-4 py-2 mr-2 border rounded-full text-slate-900 bg-slate-100 border-slate-300 w-fit">
                Student ID {student_id}
              </div>
            </a>
          </Link>
        </li>

        <ul className="flex items-center justify-end ml-auto">
          {student_id && (
            <li>
              <div className="flex items-center">
                <p className="mr-2">Amount (IDR):</p>
                <input
                  onChange={handleChangeAmount}
                  value={amount}
                  type="text"
                  className="px-4 py-2 mr-2 border rounded-full w-fit text-slate-800 hover:bg-slate-100"
                  placeholder="IDR"
                />
                <button
                  className="px-4 py-2 mr-2 text-green-900 bg-green-300 border border-green-300 rounded-full hover:bg-green-400"
                  onClick={handleAddBalance}
                >
                  {loadingButton.add ? "Loading..." : "Deposit 💰"}
                </button>
                <button
                  className="px-4 py-2 mr-2 border rounded-full border-slate-300 text-slate-800 bg-slate-300 hover:bg-slate-400"
                  onClick={handleWithdrawBalance}
                >
                  {loadingButton.withdraw ? "Loading..." : "Withdraw 🕊"}
                </button>
              </div>
            </li>
          )}
        </ul>

        {!student_id && (
          <ul className="items-center justify-between hidden max-w-xl ml-auto md:flex">
            {links.map(({ href, label }) => (
              <li key={`${href}${label}`}>
                <Link href={href}>
                  <a
                    className={`mr-4 text-gray-700 hover:text-gray-400 ${
                      router.pathname.includes(label.toLowerCase())
                        ? "font-bold"
                        : "font-normal"
                    }`}
                  >
                    {label}
                  </a>
                </Link>
              </li>
            ))}
          </ul>
        )}
        {!student_id && (
          <li className="block md:hidden">
            <BurgerMenu
              links={links}
              title={"qwe"}
              setIsOpen={setIsOpen}
              className="z-10"
            />
          </li>
        )}
      </ul>
    </nav>
  );
}

interface PropsNavItem {
  title: string;
  href: string;
}

export default Navbar;
