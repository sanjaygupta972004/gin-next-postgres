"use client";
import withAuth from "@/components/hoc/withAuth";

const Home: React.FC = () => {

  return (
    <h1 className="mt-8 text-center text-3xl font-semibold">
      Welcome to Gin + Next.js
    </h1>
  );
}


export default withAuth(Home);