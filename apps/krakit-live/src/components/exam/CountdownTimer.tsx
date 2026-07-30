import { useEffect, useState } from 'react';

type CountdownTimerProps = {
  minutes: number;
  onTimerEnd: (isFinished: boolean) => void;
};

export function CountdownTimer({ minutes, onTimerEnd }: CountdownTimerProps) {
  const [timeLeft, setTimeLeft] = useState<number>((minutes) * 60); // Convert minutes to seconds

  useEffect(() => {
    if (timeLeft <= 0) {
      onTimerEnd(true); // Notify parent when timer ends
      return;
    }

    const timer = setInterval(() => {
      setTimeLeft((prev) => prev - 1);
    }, 1000);

    return () => clearInterval(timer); // Cleanup on unmount
  }, [timeLeft, onTimerEnd]);

  // Format time as MM:SS
  const formattedTime = `${Math.floor(timeLeft / 60)
    .toString()
    .padStart(2, '0')}:${(timeLeft % 60).toString().padStart(2, '0')}`;

  return <div>{formattedTime}</div>;
}
