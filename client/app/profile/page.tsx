"use client";
import Image from "next/image";
import { useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Button } from "@/components/common/Button";
import Section from "@/components/common/Section";
import { FormInputText } from "@/components/form/textinput";
import withAuth from "@/components/hoc/withAuth";
import { isUser, User } from "@/types/user.type";
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { api_user_get_profile, api_user_update_profile, api_user_upload_avatar, api_user_upload_banner } from "@/api/user";
import { useAuth } from "@/context/AuthContext";
import { ThreeDots } from "react-loader-spinner";

const ProfilePage: React.FC = () => {
  const { setUser } = useAuth();

  const [isSaving, setIsSaving] = useState<boolean>(false);

  const hiddenBannerInput = useRef<HTMLInputElement>(null);
  const hiddenAvatarInput = useRef<HTMLInputElement>(null);
  const [bannerImage, setBannerImage] = useState<File>();
  const [avatarImage, setAvatarImage] = useState<File>();

  const getProfile = async () => {
    const me = (await api_user_get_profile()).data.data;
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
    bannerImage: yup.string()
  });

  const { control, handleSubmit, getValues, reset } = useForm<User>({
    resolver: yupResolver(userSchema), defaultValues: async () => getProfile()
  });

  const onSubmit = async (user: User) => {
    try {
      setIsSaving(true);
      if (bannerImage) {
        const imageData = new FormData();
        imageData.append('bannerImage', bannerImage);
        const res = (await api_user_upload_banner(imageData)).data.data as User;
        user.bannerImage = res.bannerImage;
      }
      if (avatarImage) {
        const imageData = new FormData();
        imageData.append('profileImage', avatarImage);
        const res = (await api_user_upload_avatar(imageData)).data.data as User;
        user.profileImage = res.profileImage;
      }
      const updated_profile = await api_user_update_profile(user);
      setUser(updated_profile.data.data as User);
    } catch (err) {
      console.error(err);
    } finally {
      setIsSaving(false);
    }
  }

  return (
    <form className="flex flex-col gap-10" onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-col relative mb-25">
        <input
          ref={hiddenBannerInput}
          type="file"
          accept="image/*"
          onChange={(e) => { if (e.target.files && e.target.files.length > 0) setBannerImage(e.target.files[0]) }}
          className="hidden"
        />
        <div className="w-full rounded-xl !z-0 min-h-[300px] max-h-[300px] bg-zinc-800 cursor-pointer overflow-hidden"
          onClick={() => hiddenBannerInput.current?.click()}
        >
          {(!!bannerImage || !!getValues("bannerImage")) &&
            <Image
              src={bannerImage ? URL.createObjectURL(bannerImage) : getValues("bannerImage")!}
              width={1000}
              height={300}
              alt="banner"
              className="w-full object-contain "
            />}
        </div>
        <input
          ref={hiddenAvatarInput}
          type="file"
          accept="image/*"
          onChange={(e) => { if (e.target.files && e.target.files.length > 0) setAvatarImage(e.target.files[0]) }}
          className="hidden"
        />
        <div
          className="absolute top-full bottom-0 left-[50%] -translate-x-[50%] -translate-y-[50%] min-h-50 min-w-50 object-contain rounded-full bg-zinc-950 p-2"
          onClick={() => hiddenAvatarInput.current?.click()}
        >
          <Image
            src={(avatarImage ? URL.createObjectURL(avatarImage) : getValues("profileImage")!) || "/avatar.png"}
            onError={(err) => { err.currentTarget.src = "/avatar.png" }}
            width={200}
            height={200}
            alt="banner"
            className="w-full h-full rounded-full cursor-pointer"
          />
        </div>
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
        <Button
          isPrimary
          customClass="w-40 flex items-center justify-center gap-2"
          disabled={isSaving}
        >
          {isSaving && <ThreeDots color="#000" width={16} height={16} />}
          Save Profile
        </Button>
        <Button isPrimary={false} customClass="w-40" onClick={() => { reset() }}>Discard</Button>
      </div>
    </form>
  );
}


export default withAuth(ProfilePage);