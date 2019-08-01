<?php
namespace KaliForms\Inc\Frontend\FormFields;

if (!defined('ABSPATH')) {
    exit;
}

/**
 * Class SubmitButton
 *
 * @package Inc\Frontend\FormFields;
 */
class SubmitButton extends Form_Field
{
    /**
     * Class constructor
     */
    public function __construct()
    {
        $this->id = 'submit-button';
    }

    /**
     * Render function
     *
     * @return void
     */
    public function render($item, $form_info)
    {
		$item['type'] = 'submit';
		$item['default'] = $item['caption'];
        $attributes = $this->generate_attribute_string($item);

        $div = '<div class="col-12 col-md-' . absint($item['col']) . '">';
        $div .= '<input ' . $attributes . '>';
        $div .= '</div>';

        return $div;

    }
}
