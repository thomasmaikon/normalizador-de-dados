import Auth from '../components/Auth'
import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import '../styles/modal.css'
import "react-toastify/dist/ReactToastify.css"
import { ToastContainer } from 'react-toastify';

function MyApp({ Component, pageProps }) {
  return (
    <Auth>
      <Component {...pageProps} />
      <ToastContainer />
    </Auth>
  )
}

export default MyApp
