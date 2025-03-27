"use client";
import * as yup from 'yup';
import { api_getme } from "@/api/auth";
import { CustomizedButton } from "@/components/common/Button";
import Section from "@/components/common/Section";
import { InputText } from "@/components/form/textinput";
import withAuth from "@/components/hoc/withAuth";
import { isUser, User } from "@/types/user.type";
import Image from "next/image";
import { useForm } from "react-hook-form";
import { yupResolver } from '@hookform/resolvers/yup';

const ProfilePage: React.FC = () => {
  const getProfile = async () => {
    const me = (await api_getme()).data;
    if (!isUser(me))
      return { name: "", email: "", role: "user" }
    else
      return me as User;
  };

  const userSchema = yup.object().shape({
    name: yup.string().required("Full name is required"),
    email: yup.string().required("Email is required"),
    role: yup.string()
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
        <InputText
          label="Name"
          control={control}
          name="name"
        />
        <InputText
          label="Email"
          control={control}
          name="email"
        />
        <InputText
          label="Role"
          control={control}
          name="role"
          disabled
        />
      </Section>
      <div className="flex justify-end items-center gap-8">
        <CustomizedButton label="Save Profile" isPrimary className="w-40" />
        <CustomizedButton label="Discard" isPrimary={false} className="w-40" />
      </div>
    </form>
  );
}


export default withAuth(ProfilePage);