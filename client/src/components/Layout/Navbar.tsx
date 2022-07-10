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
}: any): ReactElement {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const router = useRouter();
  return (
    <nav className="absolute top-0 z-10 w-full border-b">
      <ul className="flex items-center justify-between max-w-6xl px-8 py-6 mx-auto">
        <li>
          <Link href="/">
            <a className="">
              <h1 className="text-2xl font-extrabold">Kantin Kejujuran 🦋</h1>
              <h2 className="text-lg font-semibold">Balance 💸 IDR {commafy(balance)}</h2>
            </a>
          </Link>
        </li>

        <ul className="flex items-center justify-end ml-auto">
          {student_id && (
            <li>
              <div className="flex items-center">
                <button className="px-4 py-2 mr-2 font-bold text-green-900 border rounded-full bg-slate-50">
                  Student ID {student_id}
                </button>
                <button
                  className="px-4 py-2 mr-2 text-green-900 bg-green-300 rounded-full hover:bg-green-400"
                  onClick={handleAddBalance}
                >
                  {loadingButton.add ? "Loading..." : "Add IDR 5000"}
                </button>
                <button
                  className="px-4 py-2 mr-2 rounded-full text-slate-800 bg-slate-300 hover:bg-slate-400"
                  onClick={handleWithdrawBalance}
                >
                  {loadingButton.withdraw ? "Loading..." : "Withdraw IDR 5000"}
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
