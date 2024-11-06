import { AuthProvider } from "@/pkg/hook/useAuth";
import { UserProvider } from "@/pkg/hook/useUser";
import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { ToastContainer } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';

export default function App({ Component, pageProps }: AppProps) {
  return <main>
    <AuthProvider config={{redirect_to: "/app"}}>
      <UserProvider>
      <ToastContainer 
        position="top-right" 
        autoClose={5000} 
        hideProgressBar={false} 
        newestOnTop={false} 
        closeOnClick rtl={false} 
        pauseOnFocusLoss 
        draggable 
        pauseOnHover />
        <Component {...pageProps} />
      </UserProvider>
    </AuthProvider>
  </main>;
}
