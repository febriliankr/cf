import React, { ReactElement } from "react";
import Link from "next/link";
import links from "./data/links";
import { useCookies } from "react-cookie";

function Footer({ student_id }: any): ReactElement {
  const [cookies, set, removeCookie] = useCookies(["token"]);
  function handleLogout() {
    removeCookie("token");
  }

  return (
    <footer className="py-8 mx-4 lg:mx-12">
      <ul className="flex items-center justify-between py-4 md:px-8">
        <li>
          <Link href="/">
            <a className="font-bold text-gray-700">
              <div className="font-light text-gray-500">
Febrilian                © {new Date().getFullYear()}
              </div>
            </a>
          </Link>
        </li>
        {student_id && (
          <li>
            <button
              onClick={handleLogout}
              className="px-4 py-2 font-bold text-gray-800 bg-gray-300 rounded-full hover:bg-gray-400"
            >
              Logout
            </button>
          </li>
        )}
      </ul>
    </footer>
  );
}

export default Footer;
