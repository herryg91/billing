import { GetServerSideProps } from "next";

// This page means to be landing page
export default function Home() {
  return ( <div>Landing Page</div>);
}

export const getServerSideProps: GetServerSideProps = async () => {
  return {
    redirect: {
      permanent: true,
      destination: "/login",
    },
    props:{},
  };
}