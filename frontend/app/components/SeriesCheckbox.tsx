import React from "react";

const SeriesCheckbox: React.FC<{
  value: string;
  title: string;
  checked: boolean;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}> = ({ value, title, checked, onChange }) => {
  return (
    <label className="flex items-center">
      <input
        type="checkbox"
        className="form-checkbox h-6 w-6"
        onChange={onChange}
        value={value}
        checked={checked}
      />
      <span className="ml-1">{title}</span>
    </label>
  );
};

export default SeriesCheckbox;
