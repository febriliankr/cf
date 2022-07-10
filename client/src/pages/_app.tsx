import "../styles/globals.css";
import { QueryClient, QueryClientProvider } from "react-query";
import type { AppProps } from "next/app";
import { Toaster } from "react-hot-toast";
import { CookiesProvider } from "react-cookie";
const queryClient = new QueryClient();

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <CookiesProvider />
      <Toaster />
      <Component {...pageProps} />
    </QueryClientProvider>
  );
}

export default MyApp;
