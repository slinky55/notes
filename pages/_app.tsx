import { AppProps } from "next/app";
import { Toaster } from 'react-hot-toast'

import "../styles/globals.css"

export default function MyApp({ Component, pageProps }: AppProps) {
    return (
      <>
        <Toaster />
        <Component {...pageProps} />
      </>
    )
  }