import React from "react";

const SearchConditionRadio: React.FC<{
  value: string;
  title: string;
  checked: boolean;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}> = ({ value, title, checked, onChange }) => {
  return (
    <label className="flex items-center">
      <input
        type="radio"
        className="form-radio h-6 w-6"
        onChange={onChange}
        value={value}
        checked={checked}
      />
      <span className="ml-1">{title}</span>
    </label>
  );
};

export default SearchConditionRadio;
