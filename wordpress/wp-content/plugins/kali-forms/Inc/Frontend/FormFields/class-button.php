<?php
namespace KaliForms\Inc\Frontend\FormFields;

if (!defined('ABSPATH')) {
    exit;
}

/**
 * Class Button
 *
 * @package Inc\Frontend\FormFields;
 */
class Button extends Form_Field
{
    /**
     * Class constructor
     */
    public function __construct()
    {
        $this->id = 'button';
    }

    /**
     * Render function
     *
     * @return void
     */
    public function render($item, $form_info)
    {
        $item['type'] = 'button';
        $attributes = $this->generate_attribute_string($item);

        $div = '<div class="col-12 col-md-' . absint($item['col']) . '">';
        $div .= '<input ' . $attributes . '>';
        $div .= '</div>';

        return $div;

    }
}

