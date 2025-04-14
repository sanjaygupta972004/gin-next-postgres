"use client"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from 'yup';
import { FaEnvelope, FaLock, FaMale, FaUnlock, FaUser, FaUserTag } from "react-icons/fa";
import Section from "@/components/common/Section";
import { Button } from "@/components/common/Button";
import { ROUTER } from "@/constants/common";
import { useAuth } from "@/context/AuthContext";
import { ThreeDots } from "react-loader-spinner";
import { AuthRegistrationFormRequest } from "@/types/auth.type";
import { FormInputText } from "@/components/form/textinput";
import FormSelectBox from "@/components/form/selectbox";

const LoginPage: React.FC = () => {
  const { register } = useAuth();
  const router = useRouter();

  const [isRegistering, setIsRegistering] = useState<boolean>(false);

  const userSchema = yup.object().shape({
    fullName: yup.string().required('Full Name is required'),
    username: yup.string().required('Username is required'),
    email: yup.string().required('Email is required').email('Invalid email format'),
    gender: yup.string().required('Gender is required'),
    password: yup.string().required('Password')
      .min(8, 'Password must be at least 8 characters')
      .matches(
        /^(?=.*[0-9])(?=.*[!@#$%^&*])(?=.*[A-Z])/,
        'Password must contain at least one number, one special character, and one uppercase letter'
      ),
    confirmPassword: yup.string().required('Confirm password').oneOf(
      [yup.ref('password')],
      'Passwords must match'
    ),
    role: yup.string().required('Role'),
  });

  const { control, handleSubmit, formState: { errors } } = useForm<AuthRegistrationFormRequest & { confirmPassword: string }>({
    resolver: yupResolver(userSchema),
    defaultValues: { role: "user", gender: "male" }
  });

  const onSubmit = async (request: AuthRegistrationFormRequest) => {
    try {
      setIsRegistering(true);
      const user = await register(request);
      if (user) {
        router.push(ROUTER.Verification(user.userID!));
      }
    } catch (err) {
      console.error(err);
    } finally {
      setIsRegistering(false);
    }
  }

  return (
    <Section className="w-[400px] m-auto">
      <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
        <FormInputText
          label={<><FaUser />Full Name</>}
          control={control}
          name="fullName"
        />
        <FormInputText
          label={<><FaUserTag />Username</>}
          control={control}
          name="username"
        />
        <FormInputText
          label={<><FaEnvelope />Email</>}
          control={control}
          name="email"
        />
        <FormSelectBox
          label={<><FaMale />Gender</>}
          options={[
            { value: 'female', label: 'Female' },
            { value: 'male', label: 'Male' },
          ]}
          control={control}
          name="gender"
        />
        <FormInputText
          label={<><FaLock />Password</>}
          control={control}
          name="password"
          type="password"
        />
        <FormInputText
          label={<><FaUnlock />Confirm Password</>}
          control={control}
          name="confirmPassword"
          type="password"
        />
        <div className="flex justify-evenly gap-4 mt-4">
          <Button
            customClass="flex-1 flex justify-center items-center gap-2"
            type="submit"
            disabled={isRegistering}
          >
            {isRegistering && <ThreeDots color="#000" width={16} height={16} />}
            Reigster
          </Button>
          <Button
            type="button"
            customClass="flex-1 text-nowrap"
            onClick={() => router.push(ROUTER.Login)}
            isPrimary={false}
          >
            Already have?
          </Button>
        </div>
      </form>
    </Section >
  )
}

export default LoginPage;