"use client";
import Image from "next/image";
import { useForm } from "react-hook-form";
import { Button } from "@/components/common/Button";
import Section from "@/components/common/Section";
import { FormInputText } from "@/components/form/textinput";
import withAuth from "@/components/hoc/withAuth";
import { isUser, User } from "@/types/user.type";
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { api_user_profile } from "@/api/auth";

const ProfilePage: React.FC = () => {
  const getProfile = async () => {
    const me = (await api_user_profile()).data.data;
    if (!isUser(me))
      return {
        userID: "",
        fullName: "",
        gender: "",
        username: "",
        name: "",
        email: "",
        role: "user",
        profileImage: "",
        bannerImage: "",
        createdAt: Date.now().toString(),
        updatedAt: Date.now().toString()
      } as User;
    else
      return me as User;
  };

  const userSchema = yup.object().shape({
    fullName: yup.string().required("Full name is required"),
    email: yup.string().required("Email is required"),
    gender: yup.string().required("Gender is required"),
    username: yup.string().required("Username is required"),
  });

  const { control, handleSubmit } = useForm<User>({
    resolver: yupResolver(userSchema), defaultValues: async () => getProfile()
  });

  const onSubmit = (user: User) => {
    console.log(user);
  }
  return (
    <form className="flex flex-col gap-10" onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-col relative mb-25">
        <Image src="https://picsum.photos/id/20/1000/200" width={1000} height={200} alt="banner" className="w-full object-contain rounded-xl !z-0" />
        <Image
          src="https://picsum.photos/id/91/200/200"
          width={200}
          height={200}
          alt="banner"
          className="bg-zinc-950 p-2 absolute top-full bottom-0 left-[50%] -translate-x-[50%] -translate-y-[50%] w-50 object-contain rounded-full"
        />
      </div>
      <Section label="User Information">
        <FormInputText
          label="Full Name"
          control={control}
          name="fullName"
        />
        <FormInputText
          label="Username"
          control={control}
          name="username"
        />
        <FormInputText
          label="Email"
          control={control}
          name="email"
        />
        <FormInputText
          label="Role"
          control={control}
          name="role"
          disabled
        />
      </Section>
      <div className="flex justify-end items-center gap-8">
        <Button isPrimary customClass="w-40">Save Profile</Button>
        <Button isPrimary={false} customClass="w-40">Discard</Button>
      </div>
    </form>
  );
}


export default withAuth(ProfilePage);