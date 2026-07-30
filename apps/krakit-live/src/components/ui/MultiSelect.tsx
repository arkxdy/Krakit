interface Option {
  label: string;
  value: string;
  disabled?: boolean;
}

interface MultiSelectProps {
  options: Option[];
  value: string | string[];
  onChange: (value: string | string[]) => void;
  className?: string;
  mode?: "single" | "multi";
  type?: "radio" | "checkbox";
  maxSelections?: number;
}

export function MultiSelect({
  options,
  value,
  onChange,
  className = "",
  mode = "multi",
  type = "checkbox",
  maxSelections = 4,
}: MultiSelectProps) {
  const selectedValues = Array.isArray(value) ? value : [value];

  const handleOptionToggle = (optionValue: string) => {
    if (mode === "single") {
      onChange(optionValue);
    } else {
      const newValues = selectedValues.includes(optionValue)
        ? selectedValues.filter((v) => v !== optionValue)
        : [...selectedValues, optionValue];

      if (!maxSelections || newValues.length <= maxSelections) {
        onChange(newValues);
      }
    }
  };

  return (
    <div className={`space-y-2 ${className}`}>
      {options.map((option, index) => {
        const isSelected = selectedValues.includes(option.value);
        const isDisabled =
          option.disabled ||
          (mode === "multi" && maxSelections
            ? selectedValues.length >= maxSelections && !isSelected
            : false);

        return (
          <label
            key={option.value}
            className={`
              flex items-center space-x-3
              ${isDisabled ? "text-gray-500 cursor-not-allowed" : "cursor-pointer"}
            `}
          >
            <input
              type={type}
              checked={isSelected}
              onChange={() => !isDisabled && handleOptionToggle(option.value)}
              disabled={isDisabled}
              className={`
                ${type === "radio" ? "rounded-full" : "rounded"} 
                border-gray-500 text-primary-500 focus:ring-primary-500
                ${isDisabled ? "opacity-50 cursor-not-allowed" : ""}
              `}
            />
            <span className="font-semibold w-5 text-center">
              {String.fromCharCode(65 + index)}.
            </span>
            <span>{option.label}</span>
          </label>
        );
      })}
    </div>
  );
}
