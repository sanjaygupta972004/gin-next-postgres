"use client";
import Image from "next/image";

const Home: React.FC = () => {
  return (
    <div className="mt-8 text-center">
      <h1 className="text-2xl font-semibold">
        Welcome to Gin + Next.js World
      </h1>
      <Image src="/back.png" alt="logo" width={778} height={561} className="h-[480px] w-auto object-contain rounded-xl m-auto mt-16" />
      <p className="mt-16 text-lg">
        This project is designed for the Gin + Next.js community, providing a platform to share knowledge, collaborate, and build modern web applications.
      </p>
      <p className="mt-2 text-lg">
        Start building scalable, fast, and efficient web applications with this full-stack template.
      </p>

      {/* Section: About Gin */}
      <div className="mt-16 text-left">
        <h2 className="text-lg font-semibold">1. Gin</h2>
        <p className="mt-2 text-sm sm:text-base">
          Gin is a lightweight and fast web framework for Golang, designed to build high-performance RESTful APIs.
          It provides a simple yet powerful set of tools for routing, middleware, and request handling, making it ideal for modern backend development.
        </p>
        <p className="mt-2 text-sm sm:text-base">
          Learn more about Gin on the official website:{" "}
          <a
            href="https://gin-gonic.com/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-500 underline"
          >
            Gin Documentation
          </a>.
        </p>
      </div>

      {/* Section: About Next.js */}
      <div className="mt-8 text-left">
        <h2 className="text-lg font-semibold">2. Next.js</h2>
        <p className="mt-2 text-sm sm:text-base">
          Next.js is a React-based framework that enables server-side rendering, static site generation, and API routes.
          It simplifies the development of modern web applications by offering built-in features like routing, image optimization, and seamless integration with React.
        </p>
        <p className="mt-2 text-sm sm:text-base">
          Learn more about Next.js on the official website:{" "}
          <a
            href="https://nextjs.org/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-500 underline"
          >
            Next.js Documentation
          </a>.
        </p>
      </div>

      <p className="mt-16 text-lg font-semibold">{`Happy coding with Gin and Next.js. Let's get started! 🚀`}</p>
    </div>
  );
};

export default Home;