"use client"
import { useRouter, useParams } from "next/navigation";
import { useState } from "react";
import OtpInput from "react-otp-input";
import { toast } from 'react-toastify';
import { FaPaperPlane } from "react-icons/fa";
import { FiRefreshCw } from "react-icons/fi";
import Section from "@/components/common/Section";
import { Button } from "@/components/common/Button";
import { api_user_email_verify, api_user_resend_otp_code } from "@/api/auth";
import { ThreeDots } from "react-loader-spinner";
import { ROUTER } from "@/constants/common";

const LoginPage: React.FC = () => {
  const { userID } = useParams();
  const router = useRouter();

  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [isResending, setIsResending] = useState<boolean>(false);
  const [otpCode, setOtpCode] = useState<string>('');

  const onSubmit = async () => {
    try {
      setIsSubmitting(true);
      await api_user_email_verify(userID?.toString() || '', +otpCode);
      toast.success("Successfully verified!");
      router.push(ROUTER.Login);
    } catch (err) {
      console.error(err);
      // eslint-disable-next-line
      toast.error((err as any)?.message, { toastId: "otp-verification-fail" });
    } finally {
      setIsSubmitting(false);
    }
  }

  const OnResend = async () => {
    try {
      setIsResending(true);
      await api_user_resend_otp_code(userID?.toString() || '');
      setOtpCode("");
      toast.success("Successfully resent!");
    } catch (err) {
      console.error(err);
      // eslint-disable-next-line
      toast.error((err as any)?.message, { toastId: "otp-resend-fail" });
    } finally {
      setIsResending(false);
    }
  }

  return (
    <Section className="w-[400px] m-auto" label="Email OTP Verification">
      <p className="text-center">Enter the 6-digit code sent to your email</p>
      <OtpInput
        containerStyle="w-full flex justify-evenly items-center gap-2"
        inputType="number"
        value={otpCode}
        onChange={setOtpCode}
        numInputs={6}
        inputStyle="min-w-12 min-h-12 text-[22px] font-bold border border-zinc-800 hover:border-zinc-600 duration-600 rounded-[5px] focus:outline-none"
        renderInput={(props) => <input {...props} />}
      />
      <Button
        customClass="py-3 flex justify-center items-center gap-2"
        onClick={() => onSubmit()}
        disabled={isSubmitting}
      >

        {isSubmitting ? <ThreeDots color="#000" width={16} height={16} /> : <FaPaperPlane />}
        Submit
      </Button>
      <Button
        customClass="py-3 flex justify-center items-center gap-2"
        isPrimary={false}
        disabled={isResending}
        onClick={() => OnResend()}
      >
        {isResending ? <ThreeDots color="#FFF" width={16} height={16} /> : <FiRefreshCw />}
        Resend Verification Code
      </Button>
    </Section >
  )
}

export default LoginPage;